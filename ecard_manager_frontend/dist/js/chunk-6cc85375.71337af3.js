(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-6cc85375"],{6739:function(t,e,s){},"9fdb":function(t,e,s){"use strict";s("6739")},e2ad:function(t,e,s){"use strict";s.r(e);var i=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"index-img"},[s("el-row",{attrs:{gutter:20}},[s("el-col",{attrs:{span:8}},[s("el-card",{staticClass:"mgb20",staticStyle:{height:"252px"},attrs:{shadow:"hover"}},[s("div",{staticClass:"lc2-details"},[t._v("设备状态统计")]),s("g2-pie",{staticClass:"pie-chart",attrs:{type:"ring","axis-name":{name:"类别",value:"人次(次)"},data:[{name:"在线",value:this.onlineCount},{name:"离线",value:this.offlineCount}],guide:{name:"在线",value:this.onlinePercentage}}})],1),s("el-card",{staticStyle:{height:"312px"},attrs:{shadow:"hover"}},[s("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[s("span",[t._v("告警排行")])]),s("el-col",{attrs:{span:24}},[s("g2-column",{staticClass:"bar-chart",attrs:{"is-bar":!1,data:t.alarmRank,"axis-name":{name:"告警",value:"个数"},"single-color":"#1890ff"}})],1)],1)],1),s("el-col",{attrs:{span:16}},[s("el-row",{staticClass:"mgb20",attrs:{gutter:20}},[s("el-col",{attrs:{span:8}},[s("el-card",{attrs:{shadow:"hover","body-style":{padding:"0"}}},[s("div",{staticClass:"grid-content grid-con-1"},[s("i",{staticClass:"el-icon-s-custom grid-con-icon"}),s("div",{staticClass:"grid-cont-right"},[s("div",{staticClass:"grid-num"},[t._v(t._s(t.studentCount))]),t._v(" 学生总数 ")])])])],1),s("el-col",{attrs:{span:8}},[s("el-card",{attrs:{shadow:"hover","body-style":{padding:"0"}}},[s("div",{staticClass:"grid-content grid-con-2"},[s("i",{staticClass:"el-icon-mobile grid-con-icon"}),s("div",{staticClass:"grid-cont-right"},[s("div",{staticClass:"grid-num"},[t._v(t._s(t.deviceCount))]),t._v(" 设备总数 ")])])])],1),s("el-col",{attrs:{span:8}},[s("el-card",{attrs:{shadow:"hover","body-style":{padding:"0"}}},[s("div",{staticClass:"grid-content grid-con-3"},[s("i",{staticClass:"el-icon-message-solid grid-con-icon"}),s("div",{staticClass:"grid-cont-right"},[s("div",{staticClass:"grid-num"},[t._v(t._s(t.alarmCount))]),t._v(" 告警总数 ")])])])],1)],1),s("el-card",{staticStyle:{height:"463px"},attrs:{shadow:"hover"}},[s("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[s("span",[t._v("设备分布")])]),s("div",{staticStyle:{height:"463px"}},[s("amap",{attrs:{"cache-key":"coord-picker-map","mmap-style":"amap://styles/whitesmoke",async:"",center:t.center,zoom:t.zoom,"is-hotspot":""},on:{"update:center":function(e){t.center=e},"update:zoom":function(e){t.zoom=e},hotspotclick:t.onHotspotClick}},t._l(t.positions,(function(t,e){return s("div",{key:e},[s("amap-marker",{attrs:{position:t}})],1)})),0)],1)])],1)],1),s("el-dialog",{attrs:{title:"提示",visible:t.loginVisible,width:"300px","close-on-click-modal":!1,center:""},on:{"update:visible":function(e){t.loginVisible=e}}},[s("div",{staticClass:"del-dialog-cnt"},[t._v("登录时间过长，请重新登录")]),s("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[s("el-button",{on:{click:function(e){t.loginVisible=!1}}},[t._v("取 消")]),s("el-button",{attrs:{type:"primary"},on:{click:function(e){return t.Common.signOut("/login")}}},[t._v("确 定")])],1)])],1)},a=[],o={name:"dashboard",data:function(){return{Request:this.$api.api.prototype,Common:this.$common.common.prototype,loginVisible:!1,studentCount:"",deviceCount:"",onlineCount:1,offlineCount:0,onlinePercentage:"100%",alarmCount:"",loading:!0,total:0,tableData:[],searchCondition:{pageNumber:1,pageSize:4},alarmRank:[],loading1:!0,total1:0,tableData1:[],searchCondition1:{attendType:1,statusType:1,nowTime:"",pageNumber:1,pageSize:4},center:[114.4100667030411,23.10351097195999],positions:[],zoom:18}},computed:{},created:function(){this.getStudentsCountRequest(),this.getDevicesCountRequest(),this.getAlarmsCountRequest(),this.getLocationInfoRequest(),this.getAlarmRankRequest()},mounted:function(){},activated:function(){},methods:{getStudentsCountRequest:function(){this.Common.getAxios(this.Request.getStudentsCountUrl(),"",this.getStudentsCount)},getStudentsCount:function(t){if(200===t.status){var e=t.data;200===e.code&&""!==e.data&&(this.studentCount=e.data)}else 401===t.status?this.loginVisible=!0:this.$message.error("服务器繁忙，请稍后再试")},getDevicesCountRequest:function(){this.Common.getAxios(this.Request.getDevicesCountUrl(),"",this.getDevicesCount)},getDevicesCount:function(t){if(200===t.status){var e=t.data;200===e.code&&""!==e.data&&(this.deviceCount=e.data.deviceCount,this.onlineCount=parseInt(e.data.onlineCount),this.offlineCount=parseInt(e.data.offlineCount),this.onlinePercentage=Math.round(this.onlineCount/(this.onlineCount+this.offlineCount)*100)+"%")}else 401===t.status?this.loginVisible=!0:this.$message.error("服务器繁忙，请稍后再试")},getAlarmsCountRequest:function(){this.Common.getAxios(this.Request.getAlarmsCountUrl(),"",this.getAlarmsCount)},getAlarmsCount:function(t){if(200===t.status){var e=t.data;200===e.code&&""!==e.data&&(this.alarmCount=e.data)}else 401===t.status?this.loginVisible=!0:this.$message.error("服务器繁忙，请稍后再试")},getLocationInfoRequest:function(){this.Common.getAxios(this.Request.getLocationInfoUrl(),"",this.getLocationInfo)},getLocationInfo:function(t){if(200===t.status)for(var e=t.data.data,s=0;s<e.length;s++)this.positions.push([e[s].longitude,e[s].latitude]);else 401===t.status?this.loginVisible=!0:(this.$message.error("服务器繁忙，请稍后再试"),this.loading=!1)},getAlarmRankRequest:function(){this.Common.getAxios(this.Request.getAlarmRankUrl(),"",this.getAlarmRank)},getAlarmRank:function(t){200===t.status?this.alarmRank=t.data.data:401===t.status?this.loginVisible=!0:(this.$message.error("服务器繁忙，请稍后再试"),this.loading=!1)}}},n=o,l=(s("9fdb"),s("2877")),r=Object(l["a"])(n,i,a,!1,null,"348628c4",null);e["default"]=r.exports}}]);
//# sourceMappingURL=chunk-6cc85375.71337af3.js.map