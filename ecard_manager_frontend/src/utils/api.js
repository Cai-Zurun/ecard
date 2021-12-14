// const apiDomainName = 'http://1.116.213.73:8000';
const wsAddr = 'ws://1.116.213.73:3456/ws';
const apiDomainName = 'http://127.0.0.1:8000';
// const wsAddr = 'ws://127.0.0.1:3456/ws'

/* 学校信息 */
const schoolInfoUrl = apiDomainName + '/api/school/info';

/* 年级信息地址 */
const gradeInfoUrl = apiDomainName + '/pcsitesapi/api/v1/schoolInfo/grade';

/* 班级信息地址 */
const classInfoUrl = apiDomainName + '/pcsitesapi/api/v1/schoolInfo/class';

/* 学生信息地址 */
const studentUrl = apiDomainName + '/api/student';

/* 全部学生信息地址 */
const studentInfoUrl = apiDomainName + '/api/student/info';

/* 全部设备定位地址 */
const locationInfoUrl = apiDomainName + '/api/location/info';

/* 全部设备定位地址 */
const trackInfoUrl = apiDomainName + '/api/track/info';

/* 登录地址 */
const loginUrl = apiDomainName + '/api/login';

/* 获取轨迹回放地址 */
const trackListUrl =
    apiDomainName + '/position/api/v1/controllers/device/track/list';

/* 告警地址 */
const alarmUrl = apiDomainName + '/api/alarm';

/* 告警信息地址 */
const alarmInfoUrl = apiDomainName + '/api/alarm/info';

/* 告警信息地址 */
const alarmRankUrl = apiDomainName + '/api/alarm/rank';

/* 统计数量：学生 */
const studentsCountUrl = apiDomainName + '/api/student/count';

/* 统计数量：设备 */
const devicesCountUrl = apiDomainName + '/api/device/count';

/* 统计数量：告警 */
const alarmsCountUrl = apiDomainName + '/api/alarm/count';

/* 虚拟设备 */
const virtualDeviceUrl = apiDomainName + '/api/virtual-device';

/* 虚拟设备 */
const lateSituationUrl = apiDomainName + '/api/late-situation';

const fenceUrl = apiDomainName + '/api/fence';

const lateTimeUrl = apiDomainName + '/api/late-time';


class api {
  getWsAddr() {return wsAddr}

  /* 方法二：以后用这种方法统一管理api 目前先不优化了
      api不仅仅提供地址，而是要提供请求后的数据
      我在common中写有例子可提供参考put,delete,post请求自己仿造这写
  */

  /* 获取学校信息地址 */
  getSchoolInfoUrl() { return schoolInfoUrl; }

  /* 获取学校年级信息地址 */
  getGradeInfoUrl() { return gradeInfoUrl; }

  /* 获取学校班级信息地址 */
  getClassInfoUrl() { return classInfoUrl; }

  /* 获取学生信息地址 */
  getStudentUrl() { return studentUrl; }

  /* 获取全部学生信息地址 */
  getStudentInfoUrl() { return studentInfoUrl; }

  /* 获取全部设备定位地址 */
  getLocationInfoUrl() {return locationInfoUrl;}

  /* 获取全部设备定位地址 */
  getTrackInfoUrl() {return trackInfoUrl;}

  /* 获取登录信息地址 */
  getLoginUrl() { return loginUrl; }

  /* 获取轨迹回放数据地址 */
  getTrackListUrl() { return trackListUrl; }

  /* 获取告警信息地址 */
  getAlarmInfoUrl() { return alarmInfoUrl; }
  /* 告警地址 */
  getAlarmUrl() { return alarmUrl; }

  getAlarmRankUrl() { return alarmRankUrl; }

  /* 学生数量统计 */
  getStudentsCountUrl() { return studentsCountUrl; }
  /* 设备数量统计 */
  getDevicesCountUrl() { return devicesCountUrl; }
  /* 设备数量统计 */
  getAlarmsCountUrl() { return alarmsCountUrl; }

  /* 虚拟设备 */
  getVirtualDeviceUrl() {return virtualDeviceUrl; }

  /* 迟到统计 */
  getLateSituationUrl() {return lateSituationUrl; }

  getFenceUrl() {return fenceUrl; }

  getLateTimeUrl() {return lateTimeUrl; }

}

// 暴露接口
export default {api: api};
