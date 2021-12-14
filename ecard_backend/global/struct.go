package global

type LoginRequest struct {
    UserName    string  `json:"username" validate:"required"`
    Password    string  `json:"password" validate:"required"`
}

type LoginResponse struct {
    Token   string  `json:"token"`
}

type AddDeviceRequest struct {
    Imei        int      `json:"imei" validate:"required"`
}

type AddStudentRequest struct {
    Name        string  `json:"name" validate:"required"`
    StudentNo   string  `json:"studentNo" validate:"required"`
    Sex         int     `json:"sex"`
    Age         string  `json:"age" validate:"required"`
    Phone       string  `json:"phone" validate:"required"`
    Class       string  `json:"class" validate:"required"`
    Grade       string  `json:"grade" validate:"required"`
    SchoolID    int     `json:"schoolId" validate:"required"`
    Imei        string  `json:"imei" validate:"required"`
}

type GetDeviceCountResponse struct {
    DeviceCount     int `json:"deviceCount"`
    OnlineCount     int `json:"onlineCount"`
    OfflineCount    int `json:"offlineCount"`
}

type GetStudentInfoResponse struct {
    Name        string  `json:"name"`
    StudentNo   int     `json:"student_no"`
    Sex         int     `json:"sex"`
    Age         int     `json:"age"`
    Phone       string  `json:"phone"`
    Class       int     `json:"class"`
    Grade       int     `json:"grade"`
    SchoolName  string  `json:"school_name"`
    Imei        int     `json:"imei"`
}

type GetAlarmInfoResponse struct {
    AlarmTime   string  `json:"alarmTime"`
    AlarmType   string  `json:"alarmType"`
    StudentName string  `json:"name"`
    SchoolName  string  `json:"schoolName"`
    Grade       int     `json:"grade"`
    Class       int     `json:"class"`
    Phone       string  `json:"phone"`
}

type GetLocationInfoResponse struct {
    StudentName string  `json:"studentName"`
    Imei        int     `json:"imei"`
    Longitude   float64 `json:"longitude"`
    Latitude    float64 `json:"latitude"`
    LocationStr string  `json:"locationStr"`
}

type AlarmRequest struct {
    Imei        int     `json:"imei" validate:"required"`
    AlarmType   int     `json:"alarmType" validate:"required"`
}

type GetVirtualDeviceResponse struct {
    Imei        int     `json:"imei"`
    Status      int     `json:"status"`
    StudentName string  `json:"student_name"`
}

type GetTrackInfoResponse struct {
    StudentName string          `json:"studentName"`
    Imei        int             `json:"imei"`
    Track       [][2]float64    `json:"track"`
    TrackStr    []string        `json:"trackStr"`
}

type SetFenceRequest struct {
    SchoolID    int     `json:"schoolID" validate:"required"`
    Fence       string  `json:"fence" validate:"required"`
}

type GetFenceResponse struct {
    Fence   string  `json:"fence"`
}

type SetLateTimeRequest struct {
    SchoolID    int     `json:"schoolID" validate:"required"`
    LateTime    string  `json:"lateTime" validate:"required"`
}

type GetLateTimeResponse struct {
    LateTime   string  `json:"lateTime"`
}

type GetLateSituationResponse struct {
    Time        string  `json:"time"`
    StudentName string  `json:"studentName"`
    StudentNo   int     `json:"studentNo"`
    Class       int     `json:"class"`
    Grade       int     `json:"grade"`
    SchoolName  string  `json:"schoolName"`
    Phone       string  `json:"phone"`
    Situation   string  `json:"situation"`
}

type GetAlarmRankResponse struct {
    Alarm   string  `json:"name"`
    Count   int     `json:"value"`
}