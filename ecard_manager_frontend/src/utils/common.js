import axios from 'axios';
import router from '../router';
class common {
  /**
   * get 请求的公共方法
   * @param {*} url 请求地址
   * @param {*} params 请求参数（此处为params，与post，put的data不同）
   * @param {*} callback 回调函数
   * @param {*} objArray (仅在获取属性列表时使用)
   */
  getAxios(url, params, callback, objArray, obj) {
    axios({
        method: 'get',
        url: url,
        params: params
      })
      .then(res => {
        callback(res, objArray, obj);
      })
      .catch(function (error) {
        callback(error.response);
      });
  }

  /**
   * delete方法，本质与get相同（单独写只是为了调用的时候好区分，一目了然）
   * @param {*} url 
   * @param {*} params 
   * @param {*} callback 
   * @param {*} objArray 
   */
  deleteAxios(url, params, callback, objArray) {
    axios.delete(url, {
        params: params
      })
      .then(res => {
        callback(res, objArray);
      })
      .catch(function (error) {
        callback(error.response);
      });
  }

  /**
   * transformRequest 表单提交数据
   * post put 请求的公共方法（其实该方法也可以像get,delete一样分开写，此处写在一起的目的是，新增与修改提交的是同一个表单）
   * @param {*} url 请求地址
   * @param {*} type 请求类型
   * @param {*} params 请求参数
   * @param {*} callback 回调函数
   * @param {*} objArray （可存放任意你想返回的值）
   */
  commonAxios(url, type, params, callback, objArray) {
    axios({
        method: type,
        url: url,
        data: JSON.stringify(params),
      })
      .then(res => {
        callback(res, objArray);
      })
      .catch(function (error) {
        callback(error.response);
      });
  }

  /* 导出excel请求 */
  exportAxios(url, type, params, callback, objArray) {
    axios({
        method: type,
        url: url,
        data: params,
        /* 这里返回类型设置是关键 */
        responseType: 'blob',
        transformRequest: [function (data) {
          let ret = ''
          for (let it in data) {
            ret += encodeURIComponent(it) + '=' + encodeURIComponent(data[it]) + '&'
          }
          return ret
        }]
      })
      .then(res => {
        callback(res, objArray);
      })
      .catch(function (error) {
        callback(error.response);
      });
  }

  /* 导出excel信息 */
  exportExcel(data, _name) {
    const blob = new Blob([data.data], {
      type: "application/vnd.ms-excel"
    });
    const fileName = _name;
    const elink = document.createElement('a');
    elink.download = fileName;
    elink.style.display = 'none';
    elink.href = URL.createObjectURL(blob);
    document.body.appendChild(elink);
    elink.click();
    URL.revokeObjectURL(elink.href); // 释放URL 对象
    document.body.removeChild(elink);
  }

  /* 格林尼治时间转date */
  renderTime(date) {
    if (date === '') {
      return '';
    } else {
      var _date = new Date(date).toJSON();
      return new Date(+new Date(_date) + 8 * 3600 * 1000).toISOString().replace(/T/g, ' ').replace(/\.[\d]{3}Z/, '');
    }
  }

  /**
   * 跳转到空白新页面
   * @param {*} path 路径
   * @param {*} query 请求参数，是个对象
   */
  toBlankPage(path, query) {
    let routeData = router.resolve({
      path,
      query
    });
    window.open(routeData.href, '_blank');
  }
  /* 退出登录 */
  signOut(data) {
    localStorage.removeItem('token');
    router.push(data);
  }
}

/* 以后需要这么写 */
function getAxios1(url, params) {
  return new Promise((resolve, reject) => {
    axios({
        method: 'get',
        url: url,
        params: params
      })
      .then(res => {
        resolve(res.data);
      })
      .catch(function (error) {
        reject(error.data)
      });
  });
}

// 暴露接口
export default {
  common: common,
  getAxios1: getAxios1,
};