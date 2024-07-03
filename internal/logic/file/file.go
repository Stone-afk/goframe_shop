package file

import (
	"context"
	"fmt"
	"goframe_shop/internal/consts"
	"goframe_shop/internal/dao"
	"goframe_shop/internal/model"
	"goframe_shop/internal/model/entity"
	"goframe_shop/internal/service"
	"os"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type QNYServer struct {
	AccessKey string `ini:"AccessKey"`
	SecretKey string `ini:"SecretKey"`
	Bucket    string `ini:"Bucket"`
	Domain    string `ini:"Domain"`
	Zone      int    `ini:"Zone"`
}

type sFile struct{}

func init() {
	service.RegisterFile(New())
}

func New() *sFile {
	return &sFile{}
}

/*
*  Upload
1. 定义图片上传位置
2. 校验：上传位置是否正确、安全性校验：1分钟内只能上传10次
3. 定义年月日 Ymd
4. 入库
5. 返回数据
*/
func (s *sFile) Upload(ctx context.Context, in model.FileUploadInput) (out *model.FileUploadOutput, err error) {
	//定义图片上传位置
	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	if uploadPath == "" {
		return nil, gerror.New("读取配置文件失败 上传路径不存在")
	}
	if in.Name != "" {
		in.File.Filename = in.Name
	}
	count, err := dao.FileRepo.Ctx(ctx).
		Where(dao.FileRepo.Columns().UserId, gconv.Int(ctx.Value(consts.CtxAdminId))).
		WhereGTE(dao.FileRepo.Columns().CreatedAt, gtime.Now().Add(-time.Minute)).Count()
	if err != nil {
		return nil, err
	}
	//避免在代码中写死常量 抽取出去
	if count >= consts.FileMaxUploadCountMinute {
		return nil, gerror.New("上传频繁，1分钟内只能上传10次")
	}
	//	定义年月日 Ymd
	dateDirName := gtime.Now().Format("Ymd")
	//gfile.Join 用"/"拼接
	fileName, err := in.File.Save(gfile.Join(uploadPath, dateDirName), in.RandomName)
	if err != nil {
		return nil, err
	}
	//	4. 入库
	data := entity.File{
		Name:   fileName,
		Src:    gfile.Join(uploadPath, dateDirName, fileName),
		Url:    "/upload/" + dateDirName + "/" + fileName, //和上面gfile.Join()效果一样
		UserId: gconv.Int(ctx.Value(consts.CtxAdminId)),
	}
	id, err := dao.FileRepo.Ctx(ctx).Data(data).OmitEmpty().InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.FileUploadOutput{
		Id:   uint(id),
		Name: data.Name,
		Src:  data.Src,
		Url:  data.Url,
	}, nil
}

func (s *sFile) UploadCould(ctx context.Context, in model.FileUploadInput) (out *model.FileUploadOutput, err error) {
	qNYServer := LoadQiniuCfg(ctx)

	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	if uploadPath == "" {
		return nil, gerror.New("读取配置文件失败 上传路径不存在")
	}
	if in.Name != "" {
		in.File.Filename = in.Name
	}
	//	定义年月日 Ymd
	dateDirName := gtime.Now().Format("Ymd")
	//gfile.Join 用"/"拼接
	dirPath := gfile.Join(uploadPath, dateDirName)
	fileName, err := in.File.Save(dirPath, in.RandomName)
	if err != nil {
		return nil, err
	}
	//	定义本地文件路径
	localFile := dirPath + fileName
	putPolicy := storage.PutPolicy{
		Scope: qNYServer.Bucket,
	}

	// 获取上传token
	mac := qbox.NewMac(qNYServer.AccessKey, qNYServer.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	// 设置上传配置
	// 构建表单上传的对象
	// 根据自己需求去灵活配置
	formUploader := storage.NewFormUploader(&storage.Config{
		Region:        selectZone(qNYServer.Zone),
		UseHTTPS:      true,
		UseCdnDomains: false,
	})
	//上传结果的结构体
	ret := storage.PutRet{}
	//可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	//七牛云表单上传
	key := fileName
	err = formUploader.PutFile(ctx, &ret, upToken, key, localFile, &putExtra)
	g.Dump(err)
	if err != nil {
		return nil, err
	}
	fmt.Println(ret.Key, ret.Hash, ret.PersistentID)
	//删除本地临时文件
	_ = os.RemoveAll(localFile)
	//if err != nil {
	//	return nil, err
	//}
	// 返回数据
	url := qNYServer.Domain + ret.Key
	return &model.FileUploadOutput{
		Name: in.Name,
		Src:  qNYServer.Domain,
		Url:  url,
	}, nil
}

func selectZone(Zone int) *storage.Zone {
	switch Zone {
	case 1:
		return &storage.ZoneHuadong
	case 2:
		return &storage.ZoneHuadongZheJiang2
	case 3:
		return &storage.ZoneHuabei
	case 4:
		return &storage.ZoneHuanan
	default:
		return &storage.ZoneHuadong
	}
}

func LoadQiniuCfg(ctx context.Context) *QNYServer {
	return &QNYServer{
		Zone:      g.Cfg().MustGet(ctx, "qiniuyun.zone").Int(),
		Bucket:    g.Cfg().MustGet(ctx, "qiniuyun.bucket").String(),
		AccessKey: g.Cfg().MustGet(ctx, "qiniuyun.accessKey").String(),
		SecretKey: g.Cfg().MustGet(ctx, "qiniuyun.secretKey").String(),
		Domain:    g.Cfg().MustGet(ctx, "qiniuyun.domain").String(),
	}
}
