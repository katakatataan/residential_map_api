package controller

// ここのコントローラーはその外側のレイヤーから呼び出される
// interfaceもstructに対して適応される
//　ここにいぞんさせて実際にnewさせるのはinterfaceを実装したコントローラー
type Controller interface {
	GetMstPrefecture()
}

type MstPrefCityController struct {
	Intaractor MstPrefCityInteractor
}

func NewMstPrefCityController() *MstPrefCityController {
	return &MstPrefCityController{}
}

func GetMstPrefCity(mpc *MstPrefCityController) {
	//ここで実際の取得処理を書く
	hc.Intaractor.get
}
