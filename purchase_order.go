package determine

import (
	"encoding/xml"

	"github.com/omniboost/go-determine/types"
)

type BPACK struct {
	XMLName xml.Name `xml:"BPACK"`

	Xsi      string `xml:"xsi,attr"`
	RecordID string `xml:"recordid,attr"`
	Format   string `xml:"format,attr"`
	Version  string `xml:"version,attr"`
	Lang     string `xml:"lang,attr"`
	Order    Order  `xml:"record.ORDERID"`
}

type Order struct {
	XMLName xml.Name `xml:"record.ORDERID"`

	// Text        string `xml:",chardata"`
	DocNum             string         `xml:"DOCNUM"`
	Description        string         `xml:"DESCRIPTION"`
	StampDate          string         `xml:"STAMPDATE"`
	SupplierID         ID             `xml:"SUPPLIERID"`
	Status             Status         `xml:"STATUS"`
	UserID             ID             `xml:"USERID"`
	InvCompanyID       ID             `xml:"INVCOMPANYID"`
	InvDPTID           DPTID          `xml:"INVDPTID"`
	Priority           Priority       `xml:"PRIORITY"`
	NEEDEDDATE         string         `xml:"NEEDEDDATE"`
	Info               Info           `xml:"INFO"`
	RecDPTID           DPTID          `xml:"RECDPTID"`
	RecDate            Date           `xml:"RECDATE"`
	RequestIDs         IDs            `xml:"REQUESTIDS"`
	ClientID           ID             `xml:"CLIENTID"`
	SupContactID       ID             `xml:"SUPCONTACTID"`
	AckDate            Date           `xml:"ACKDATE"`
	EvalDate           Date           `xml:"EVALDATE"`
	PrevDate           Date           `xml:"PREVDATE"`
	DoneDate           Date           `xml:"DONEDATE"`
	DueDate            Date           `xml:"DUEDATE"`
	UStampDate         types.DateTime `xml:"USTAMPDATE"`
	ReqType            string         `xml:"REQTYPE"`
	OrderType          string         `xml:"ORDERTYPE"`
	CurID              ID             `xml:"CURID"`
	EditDate           Date           `xml:"EDITDATE"`
	PaieCode           Code           `xml:"PAIECODE"`
	TotalHT            string         `xml:"TOTALHT"`
	ICURIDTOTALHT      string         `xml:"ICURIDTOTALHT"`
	TotalTTC           string         `xml:"TOTALTTC"`
	CTOTALHT           string         `xml:"CTOTALHT"`
	CTOTALTTC          string         `xml:"CTOTALTTC"`
	AStatus            Status         `xml:"ASTATUS"`
	OrderID            string         `xml:"ORDERID"`
	IStatus            Status         `xml:"ISTATUS"`
	CopyFromID         ID             `xml:"COPYFROMID"`
	INVCOMPGRP         String         `xml:"INVCOMPGRP"`
	NEWAMOUNT          Amount         `xml:"NEWAMOUNT"`
	Temp2Load          String         `xml:"Temp2Load"`
	NoSupEdit          String         `xml:"NoSupEdit"`
	DateVal            string         `xml:"DateVal"`
	SENTDATE           Date           `xml:"SENTDATE"`
	ATTR               String         `xml:"ATTR"`
	XATTR              String         `xml:"XATTR"`
	ORDERPARAM         String         `xml:"ORDERPARAM"`
	JSLISTENER         String         `xml:"JS_LISTENER"`
	LASTAPPROBID       ID             `xml:"LASTAPPROBID"`
	LASTAPPROVALDATE   Date           `xml:"LASTAPPROVALDATE"`
	USERADDR           String         `xml:"USERADDR"`
	DocAttr            String         `xml:"DocAttr"`
	REQUSERID          ID             `xml:"REQUSERID"`
	DISCOUNT           Amount         `xml:"DISCOUNT"`
	INITIALAMOUNT      Amount         `xml:"INITIALAMOUNT"`
	RECUSERID          ID             `xml:"RECUSERID"`
	EDITUSERID         ID             `xml:"EDITUSERID"`
	VALIDATED          Bool           `xml:"VALIDATED"`
	BUYERMODE          String         `xml:"BUYERMODE"`
	PCHECK             String         `xml:"PCHECK"`
	ACCOUNTNB          String         `xml:"ACCOUNTNB"`
	SUPCOMMENT         String         `xml:"SUPCOMMENT"`
	CONTRACTCODE       String         `xml:"CONTRACTCODE"`
	WORKER             String         `xml:"WORKER"`
	EDITDOC            String         `xml:"EDITDOC"`
	OPCODE             String         `xml:"OPCODE"`
	INVESTCODE         String         `xml:"INVESTCODE"`
	SIGNID             ID             `xml:"SIGNID"`
	ICURID             ID             `xml:"ICURID"`
	ICURIDRATE         string         `xml:"ICURIDRATE"`
	AMOUNTVAT          Amount         `xml:"AMOUNTVAT"`
	HASFOLLOW          Bool           `xml:"HASFOLLOW"`
	NBLINEITEMS        String         `xml:"NBLINEITEMS"`
	PORECDELAY         String         `xml:"PO_RECDELAY"`
	REQUESTID          ID             `xml:"REQUESTID"`
	INVFILEID          ID             `xml:"INVFILEID"`
	CATID              ID             `xml:"CATID"`
	CATIDSUB           String         `xml:"CATID_SUB"`
	CCURID             ID             `xml:"CCURID"`
	WAPPROVER          String         `xml:"WAPPROVER"`
	RECORDNAME         string         `xml:"RECORDNAME"`
	RECORDZOOM         String         `xml:"RECORDZOOM"`
	USERIDCONFIGCHAR3  String         `xml:"USERID_CONFIGCHAR3"`
	INVCOMPANYIDCODE   string         `xml:"INVCOMPANYID_CODE"`
	SUPPLIERIDCODE     string         `xml:"SUPPLIERID_CODE"`
	USERIDNAME         string         `xml:"USERID_NAME"`
	RECDPTIDCODE       string         `xml:"RECDPTID_CODE"`
	CURIDDATE          string         `xml:"CURIDDATE"`
	OrderMode          String         `xml:"OrderMode"`
	MMSendTo           String         `xml:"MM_SendTo"`
	LASTACTIONAPPROBID ID             `xml:"LASTACTIONAPPROBID"`
	CURRENTAPPROBID    ID             `xml:"CURRENTAPPROBID"`
	INITUSERID         string         `xml:"INITUSERID"`
	MODUSERID          ID             `xml:"MODUSERID"`
	LASTUPDATED        string         `xml:"LASTUPDATED"`
	DOCAGE             String         `xml:"DOC_AGE"`
	ERPDATE            Date           `xml:"ERPDATE"`
	ERPFLAG            Bool           `xml:"ERPFLAG"`
	ERPCODE            String         `xml:"ERPCODE"`
	ERPERROR           String         `xml:"ERPERROR"`
	APPROVDATE         string         `xml:"APPROVDATE"`
	DOCSTATUS          String         `xml:"DOCSTATUS"`
	DOCISTATUS         Status         `xml:"DOCISTATUS"`
	INMODIF            String         `xml:"INMODIF"`
	ASENTDATE          string         `xml:"ASENTDATE"`
	LASTAPPROBAMOUNT   Amount         `xml:"LASTAPPROBAMOUNT"`
	Subrecords         []SubRecord    `xml:"subrecords"`
}

type SubRecord struct {
	RecordDOCITEMID       []DocItemID      `xml:"record.DOCITEMID"`
	RecordCOSTALLOCATIONS []CostAllocation `xml:"record.COSTALLOCATIONS"`
	RecordACCOUNTS        []Account        `xml:"record.ACCOUNTS"`
}

type DocItemID struct {
	ITEMID          ID     `xml:"ITEMID"`
	ANAME           string `xml:"ANAME"`
	SUPPLIERID      ID     `xml:"SUPPLIERID"`
	SUPREF          String `xml:"SUPREF"`
	SUPPARTAUXID    ID     `xml:"SUPPARTAUXID"`
	CLIENTID        ID     `xml:"CLIENTID"`
	SHIPDELAY       String `xml:"SHIPDELAY"`
	CATID           string `xml:"CATID"`
	ITYPE           String `xml:"ITYPE"`
	ADESCRIPTION    String `xml:"ADESCRIPTION"`
	THUMBID         ID     `xml:"THUMBID"`
	AREF            String `xml:"AREF"`
	MANUFACTURER    String `xml:"MANUFACTURER"`
	MANREF          String `xml:"MANREF"`
	ITEMQUANTITY    string `xml:"ITEMQUANTITY"`
	QTY             string `xml:"QTY"`
	CODEUNIT        string `xml:"CODEUNIT"`
	PUHT            string `xml:"PUHT"`
	CURID           string `xml:"CURID"`
	USRREBATE       String `xml:"USRREBATE"`
	ARTREBATE       String `xml:"ARTREBATE"`
	REBATE          String `xml:"REBATE"`
	CODETVA         string `xml:"CODETVA"`
	TVA             string `xml:"TVA"`
	INVDPTID        ID     `xml:"INVDPTID"`
	INVEST          string `xml:"INVEST"`
	ACCOUNTID       string `xml:"ACCOUNTID"`
	INVCOMPGRP      string `xml:"INVCOMPGRP"`
	STARTDATE       Date   `xml:"STARTDATE"`
	ENDDATE         Date   `xml:"ENDDATE"`
	DateVal         String `xml:"DateVal"`
	CTRDATE         Date   `xml:"CTRDATE"`
	REQUESTID       ID     `xml:"REQUESTID"`
	ITEMNUMBER      string `xml:"ITEMNUMBER"`
	ORDERID         ID     `xml:"ORDERID"`
	ORDERISTATUS    string `xml:"ORDER_ISTATUS"`
	AFROM           String `xml:"AFROM"`
	OLDSUPPLIERID   ID     `xml:"OLDSUPPLIERID"`
	INVQTY          Amount `xml:"INVQTY"`
	INVPUHT         String `xml:"INVPUHT"`
	RECQTY          String `xml:"RECQTY"`
	RECDATE         Date   `xml:"RECDATE"`
	CURCONV         string `xml:"CURCONV"`
	CURDATE         string `xml:"CURDATE"`
	CTOTALHT        string `xml:"CTOTALHT"`
	CTOTALREF       String `xml:"CTOTALREF"`
	CPUHT           string `xml:"CPUHT"`
	PUNET           string `xml:"PUNET"`
	CPUNET          string `xml:"CPUNET"`
	AMOUNTVAT       Amount `xml:"AMOUNTVAT"`
	ICURIDTOTALHT   string `xml:"ICURIDTOTALHT"`
	ICURIDAMOUNTVAT Amount `xml:"ICURIDAMOUNTVAT"`
	CAMOUNTVAT      Amount `xml:"CAMOUNTVAT"`
	ACCOUNTTVA      String `xml:"ACCOUNTTVA"`
	ACCOUNTTTC      String `xml:"ACCOUNTTTC"`
	NEEDEDDATE      string `xml:"NEEDEDDATE"`
	REFITYPE        String `xml:"REF_ITYPE"`
	SUMITYPE        String `xml:"SUM_ITYPE"`
	INVCOMPANYID    ID     `xml:"INVCOMPANYID"`
	ACCGROUPID      ID     `xml:"ACCGROUPID"`
	DATEVAL         String `xml:"DATEVAL"`
	PREVDATE        string `xml:"PREVDATE"`
	TOTALHT         string `xml:"TOTALHT"`
	PACKID          ID     `xml:"PACKID"`
	CREDITQTY       String `xml:"CREDITQTY"`
	USTAMPDATE      string `xml:"USTAMPDATE"`
	DELETED         Bool   `xml:"DELETED"`
	MODDATE         Date   `xml:"MODDATE"`
	DELDATE         Date   `xml:"DELDATE"`
	REFITEMID       ID     `xml:"REFITEMID"`
	COMMENTS        String `xml:"COMMENTS"`
	ARTICLEID       ID     `xml:"ARTICLEID"`
	PRICEID         string `xml:"PRICEID"`
	REQTYPE         string `xml:"REQTYPE"`
	ACTION          String `xml:"ACTION"`
	ACTIONTYPE      String `xml:"ACTION_TYPE"`
	COUNTRYID       ID     `xml:"COUNTRYID"`
	TEMPCATTRITEMID ID     `xml:"TEMP_CATTRITEMID"`
	CCURID          ID     `xml:"CCURID"`
	ERPFLAG         Bool   `xml:"ERPFLAG"`
	TAGIDS          IDs    `xml:"TAGIDS"`
	CURIDDATE       string `xml:"CURIDDATE"`
	ICURID          string `xml:"ICURID"`
	ICURIDRATE      string `xml:"ICURIDRATE"`
	ERPDATE         Date   `xml:"ERPDATE"`
	INVDPTIDCODE    string `xml:"INVDPTID_CODE"`
}

type CostAllocation struct {
	ITEMID          ID     `xml:"ITEMID"`
	ANAME           string `xml:"ANAME"`
	SUPPLIERID      ID     `xml:"SUPPLIERID"`
	SUPREF          String `xml:"SUPREF"`
	SUPPARTAUXID    ID     `xml:"SUPPARTAUXID"`
	CLIENTID        ID     `xml:"CLIENTID"`
	SHIPDELAY       String `xml:"SHIPDELAY"`
	CATID           string `xml:"CATID"`
	ITYPE           String `xml:"ITYPE"`
	ADESCRIPTION    String `xml:"ADESCRIPTION"`
	THUMBID         ID     `xml:"THUMBID"`
	AREF            String `xml:"AREF"`
	MANUFACTURER    String `xml:"MANUFACTURER"`
	MANREF          String `xml:"MANREF"`
	ITEMQUANTITY    string `xml:"ITEMQUANTITY"`
	QTY             string `xml:"QTY"`
	CODEUNIT        string `xml:"CODEUNIT"`
	PUHT            string `xml:"PUHT"`
	CURID           string `xml:"CURID"`
	USRREBATE       String `xml:"USRREBATE"`
	ARTREBATE       String `xml:"ARTREBATE"`
	REBATE          String `xml:"REBATE"`
	CODETVA         string `xml:"CODETVA"`
	TVA             string `xml:"TVA"`
	INVDPTID        DPTID  `xml:"INVDPTID"`
	INVEST          string `xml:"INVEST"`
	ACCOUNTID       string `xml:"ACCOUNTID"`
	INVCOMPGRP      string `xml:"INVCOMPGRP"`
	STARTDATE       Date   `xml:"STARTDATE"`
	ENDDATE         Date   `xml:"ENDDATE"`
	DateVal         String `xml:"DateVal"`
	CTRDATE         Date   `xml:"CTRDATE"`
	REQUESTID       ID     `xml:"REQUESTID"`
	ITEMNUMBER      string `xml:"ITEMNUMBER"`
	ORDERID         ID     `xml:"ORDERID"`
	ORDERISTATUS    string `xml:"ORDER_ISTATUS"`
	AFROM           String `xml:"AFROM"`
	OLDSUPPLIERID   ID     `xml:"OLDSUPPLIERID"`
	INVQTY          Amount `xml:"INVQTY"`
	INVPUHT         String `xml:"INVPUHT"`
	RECQTY          String `xml:"RECQTY"`
	RECDATE         Date   `xml:"RECDATE"`
	CURCONV         string `xml:"CURCONV"`
	CURDATE         string `xml:"CURDATE"`
	CTOTALHT        string `xml:"CTOTALHT"`
	CTOTALREF       String `xml:"CTOTALREF"`
	CPUHT           string `xml:"CPUHT"`
	PUNET           string `xml:"PUNET"`
	CPUNET          string `xml:"CPUNET"`
	AMOUNTVAT       Amount `xml:"AMOUNTVAT"`
	ICURIDTOTALHT   string `xml:"ICURIDTOTALHT"`
	ICURIDAMOUNTVAT Amount `xml:"ICURIDAMOUNTVAT"`
	CAMOUNTVAT      Amount `xml:"CAMOUNTVAT"`
	ACCOUNTTVA      String `xml:"ACCOUNTTVA"`
	ACCOUNTTTC      String `xml:"ACCOUNTTTC"`
	NEEDEDDATE      Date   `xml:"NEEDEDDATE"`
	REFITYPE        String `xml:"REF_ITYPE"`
	SUMITYPE        String `xml:"SUM_ITYPE"`
	INVCOMPANYID    ID     `xml:"INVCOMPANYID"`
	ACCGROUPID      ID     `xml:"ACCGROUPID"`
	DATEVAL         String `xml:"DATEVAL"`
	PREVDATE        string `xml:"PREVDATE"`
	TOTALHT         string `xml:"TOTALHT"`
	PACKID          ID     `xml:"PACKID"`
	CREDITQTY       Amount `xml:"CREDITQTY"`
	USTAMPDATE      string `xml:"USTAMPDATE"`
	DELETED         Bool   `xml:"DELETED"`
	MODDATE         Date   `xml:"MODDATE"`
	DELDATE         Date   `xml:"DELDATE"`
	REFITEMID       ID     `xml:"REFITEMID"`
	COMMENTS        String `xml:"COMMENTS"`
	ARTICLEID       ID     `xml:"ARTICLEID"`
	PRICEID         string `xml:"PRICEID"`
	REQTYPE         string `xml:"REQTYPE"`
	ACTION          String `xml:"ACTION"`
	COUNTRYID       ID     `xml:"COUNTRYID"`
	TEMPCATTRITEMID ID     `xml:"TEMP_CATTRITEMID"`
	CCURID          ID     `xml:"CCURID"`
	ERPFLAG         String `xml:"ERPFLAG"`
	TAGIDS          IDs    `xml:"TAGIDS"`
	CURIDDATE       string `xml:"CURIDDATE"`
	ICURID          string `xml:"ICURID"`
	ICURIDRATE      string `xml:"ICURIDRATE"`
	ERPDATE         Date   `xml:"ERPDATE"`
}

type Account struct {
	ITEMID          string `xml:"ITEMID"`
	ANAME           string `xml:"ANAME"`
	SUPPLIERID      ID     `xml:"SUPPLIERID"`
	SUPREF          String `xml:"SUPREF"`
	SUPPARTAUXID    ID     `xml:"SUPPARTAUXID"`
	CLIENTID        ID     `xml:"CLIENTID"`
	SHIPDELAY       String `xml:"SHIPDELAY"`
	CATID           string `xml:"CATID"`
	ITYPE           String `xml:"ITYPE"`
	ADESCRIPTION    String `xml:"ADESCRIPTION"`
	THUMBID         ID     `xml:"THUMBID"`
	AREF            String `xml:"AREF"`
	MANUFACTURER    String `xml:"MANUFACTURER"`
	MANREF          String `xml:"MANREF"`
	ITEMQUANTITY    string `xml:"ITEMQUANTITY"`
	QTY             string `xml:"QTY"`
	CODEUNIT        string `xml:"CODEUNIT"`
	PUHT            string `xml:"PUHT"`
	CURID           string `xml:"CURID"`
	USRREBATE       String `xml:"USRREBATE"`
	ARTREBATE       String `xml:"ARTREBATE"`
	REBATE          String `xml:"REBATE"`
	CODETVA         string `xml:"CODETVA"`
	TVA             string `xml:"TVA"`
	INVDPTID        ID     `xml:"INVDPTID"`
	INVEST          string `xml:"INVEST"`
	ACCOUNTID       string `xml:"ACCOUNTID"`
	INVCOMPGRP      string `xml:"INVCOMPGRP"`
	STARTDATE       Date   `xml:"STARTDATE"`
	ENDDATE         Date   `xml:"ENDDATE"`
	DateVal         String `xml:"DateVal"`
	CTRDATE         Date   `xml:"CTRDATE"`
	REQUESTID       ID     `xml:"REQUESTID"`
	ITEMNUMBER      string `xml:"ITEMNUMBER"`
	ORDERID         ID     `xml:"ORDERID"`
	ORDERISTATUS    string `xml:"ORDER_ISTATUS"`
	AFROM           String `xml:"AFROM"`
	OLDSUPPLIERID   ID     `xml:"OLDSUPPLIERID"`
	INVQTY          Amount `xml:"INVQTY"`
	INVPUHT         String `xml:"INVPUHT"`
	RECQTY          Amount `xml:"RECQTY"`
	RECDATE         Date   `xml:"RECDATE"`
	CURCONV         string `xml:"CURCONV"`
	CURDATE         string `xml:"CURDATE"`
	CTOTALHT        string `xml:"CTOTALHT"`
	CTOTALREF       String `xml:"CTOTALREF"`
	CPUHT           string `xml:"CPUHT"`
	PUNET           string `xml:"PUNET"`
	CPUNET          string `xml:"CPUNET"`
	AMOUNTVAT       Amount `xml:"AMOUNTVAT"`
	ICURIDTOTALHT   string `xml:"ICURIDTOTALHT"`
	ICURIDAMOUNTVAT Amount `xml:"ICURIDAMOUNTVAT"`
	CAMOUNTVAT      Amount `xml:"CAMOUNTVAT"`
	ACCOUNTTVA      String `xml:"ACCOUNTTVA"`
	ACCOUNTTTC      String `xml:"ACCOUNTTTC"`
	NEEDEDDATE      string `xml:"NEEDEDDATE"`
	REFITYPE        String `xml:"REF_ITYPE"`
	SUMITYPE        String `xml:"SUM_ITYPE"`
	INVCOMPANYID    ID     `xml:"INVCOMPANYID"`
	ACCGROUPID      ID     `xml:"ACCGROUPID"`
	DATEVAL         String `xml:"DATEVAL"`
	PREVDATE        string `xml:"PREVDATE"`
	TOTALHT         string `xml:"TOTALHT"`
	PACKID          ID     `xml:"PACKID"`
	CREDITQTY       Amount `xml:"CREDITQTY"`
	USTAMPDATE      Date   `xml:"USTAMPDATE"`
	DELETED         Bool   `xml:"DELETED"`
	MODDATE         Date   `xml:"MODDATE"`
	DELDATE         Date   `xml:"DELDATE"`
	REFITEMID       ID     `xml:"REFITEMID"`
	COMMENTS        String `xml:"COMMENTS"`
	ARTICLEID       ID     `xml:"ARTICLEID"`
	PRICEID         string `xml:"PRICEID"`
	REQTYPE         string `xml:"REQTYPE"`
	ACTION          String `xml:"ACTION"`
	COUNTRYID       ID     `xml:"COUNTRYID"`
	TEMPCATTRITEMID ID     `xml:"TEMP_CATTRITEMID"`
	CCURID          ID     `xml:"CCURID"`
	ERPFLAG         String `xml:"ERPFLAG"`
	TAGIDS          IDs    `xml:"TAGIDS"`
	CURIDDATE       Date   `xml:"CURIDDATE"`
	ICURID          ID     `xml:"ICURID"`
	ICURIDRATE      string `xml:"ICURIDRATE"`
	ERPDATE         Date   `xml:"ERPDATE"`
}

type ID struct {
	ID    string `xml:",chardata"`
	Label string `xml:"label,attr"`
	Nil   bool   `xml:"nil,attr,omitempty"`
}

type Status struct {
	ID    string `xml:",chardata"`
	Label string `xml:"label,attr"`
}

type DPTID struct {
	Text  string `xml:",chardata"`
	Nil   bool   `xml:"nil,attr,omitempty"`
	Label string `xml:"label,attr"`
}

type Priority struct {
	ID    string `xml:",chardata"`
	Label string `xml:"label,attr"`
}

type Info struct {
	Text string `xml:",chardata"`
	Nil  bool   `xml:"nil,attr,omitempty"`
}

type Date struct {
	Date types.Date `xml:",chardata"`
	Nil  bool       `xml:"nil,attr,omitempty"`
}

type Code struct {
	Code  string `xml:",chardata"`
	Label string `xml:"label,attr"`
}

type IDs []string

type String struct {
	Text string `xml:",chardata"`
	Nil  string `xml:"nil,attr,omitempty"`
}

type Bool struct {
	Value bool   `xml:",chardata"`
	Label string `xml:"label,attr,omitempty"`
}

type Amount struct {
	Text string `xml:",chardata"`
	Nil  string `xml:"nil,attr,omitempty"`
}
