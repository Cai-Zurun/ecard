(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-11c48044"],{1911:function(t,e,l){"use strict";l.r(e);var s=function(){var t=this,e=t.$createElement,l=t._self._c||e;return l("div",{staticClass:"table"},[l("div",{staticClass:"container"},[l("el-row",{staticClass:"mgb20",staticStyle:{"margin-top":"10px"},attrs:{gutter:20}},[l("el-col",{attrs:{span:8}},[l("el-button",{attrs:{type:"primary",icon:"el-icon-plus"},on:{click:t.addVirtualDevice}},[t._v(" 添加虚拟设备 ")])],1)],1),l("el-divider"),l("el-row",{staticClass:"mgb20",attrs:{gutter:30}},t._l(t.virtualDevices,(function(e,s){return l("el-col",{key:s,attrs:{span:8}},[l("el-card",{attrs:{shadow:"always","body-style":{padding:"0"}}},[l("div",{staticClass:"grid-content grid-con-2"},[l("i",{staticClass:"el-icon-mobile-phone grid-con-icon"}),l("div",{staticClass:"grid-cont-right"},[l("div",{staticStyle:{"margin-bottom":"5px"}},[t._v("imei: "+t._s(e.imei))]),l("div",{staticStyle:{"margin-bottom":"15px"}},[t._v("学生姓名: "+t._s(e.student_name))]),l("el-button",{attrs:{type:"danger",icon:"el-icon-bell"},on:{click:function(l){return t.alarmRequest(e.imei)}}},[t._v(" 告警 ")])],1)])])],1)})),1)],1),l("el-dialog",{attrs:{title:"提示",visible:t.loginVisible,width:"300px","close-on-click-modal":!1,center:""},on:{"update:visible":function(e){t.loginVisible=e}}},[l("div",{staticClass:"del-dialog-cnt"},[t._v("登录时间过长，请重新登录")]),l("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[l("el-button",{on:{click:function(e){t.loginVisible=!1}}},[t._v("取 消")]),l("el-button",{attrs:{type:"primary"},on:{click:function(e){return t.Common.signOut("/login")}}},[t._v("确 定")])],1)]),l("el-dialog",{attrs:{title:t.dialogTitle,visible:t.editVisible,width:"40%","close-on-click-modal":!1},on:{"update:visible":function(e){t.editVisible=e},closed:t.closedFormDialog}},[l("el-form",{directives:[{name:"loading",rawName:"v-loading",value:t.loadingDialog,expression:"loadingDialog"}],ref:"imeiForm",attrs:{model:t.studentForm,rules:t.rules,"label-width":"120px"}},[l("el-row",{attrs:{gutter:20}},[l("el-col",{attrs:{span:20}},[l("el-form-item",{attrs:{label:"学号：",prop:"studentNo"}},[l("el-input",{attrs:{placeholder:"请输入学号"},model:{value:t.studentForm.studentNo,callback:function(e){t.$set(t.studentForm,"studentNo",e)},expression:"studentForm.studentNo"}})],1)],1)],1),l("el-row",{attrs:{gutter:20}},[l("el-col",{attrs:{span:20}},[l("el-form-item",{attrs:{label:"姓名：",prop:"name"}},[l("el-input",{attrs:{placeholder:"请输入姓名"},model:{value:t.studentForm.name,callback:function(e){t.$set(t.studentForm,"name",e)},expression:"studentForm.name"}})],1)],1)],1),l("el-row",{attrs:{gutter:20}},[l("el-col",{attrs:{span:20}},[l("el-form-item",{attrs:{label:"年龄：",prop:"age"}},[l("el-input",{attrs:{placeholder:"请输入年龄"},model:{value:t.studentForm.age,callback:function(e){t.$set(t.studentForm,"age",e)},expression:"studentForm.age"}})],1)],1)],1),l("el-row",{attrs:{gutter:20}},[l("el-form-item",{attrs:{label:"性别：",prop:"sex"}},[l("el-radio-group",{model:{value:t.studentForm.sex,callback:function(e){t.$set(t.studentForm,"sex",e)},expression:"studentForm.sex"}},[l("el-radio",{attrs:{label:1}},[t._v("男")]),l("el-radio",{attrs:{label:0}},[t._v("女")])],1)],1)],1),l("el-row",{attrs:{gutter:20}},[l("el-col",{attrs:{span:20}},[l("el-form-item",{attrs:{label:"手机：",prop:"phone"}},[l("el-input",{attrs:{placeholder:"请输入手机"},model:{value:t.studentForm.phone,callback:function(e){t.$set(t.studentForm,"phone",e)},expression:"studentForm.phone"}})],1)],1)],1),l("el-row",{attrs:{gutter:20}},[l("el-col",{attrs:{span:20}},[l("el-form-item",{attrs:{label:"学校：",prop:"schoolId"}},[l("el-select",{attrs:{placeholder:"请选择学校"},model:{value:t.studentForm.schoolId,callback:function(e){t.$set(t.studentForm,"schoolId",e)},expression:"studentForm.schoolId"}},t._l(t.selectSchool,(function(t){return l("el-option",{key:t.ID,attrs:{label:t.SchoolName,value:t.ID}})})),1)],1)],1)],1),l("el-row",{attrs:{gutter:20}},[l("el-col",{attrs:{span:20}},[l("el-form-item",{attrs:{label:"年级：",prop:"grade"}},[l("el-input",{attrs:{placeholder:"请输入年级"},model:{value:t.studentForm.grade,callback:function(e){t.$set(t.studentForm,"grade",e)},expression:"studentForm.grade"}})],1)],1)],1),l("el-row",{attrs:{gutter:20}},[l("el-col",{attrs:{span:20}},[l("el-form-item",{attrs:{label:"班级：",prop:"class"}},[l("el-input",{attrs:{placeholder:"请输入班级"},model:{value:t.studentForm.class,callback:function(e){t.$set(t.studentForm,"class",e)},expression:"studentForm.class"}})],1)],1)],1),l("el-row",{attrs:{gutter:20}},[l("el-col",{attrs:{span:20}},[l("el-form-item",{attrs:{label:"IMEI号：",prop:"imei"}},[l("el-input",{attrs:{placeholder:"请输入IMEI号"},model:{value:t.studentForm.imei,callback:function(e){t.$set(t.studentForm,"imei",e)},expression:"studentForm.imei"}})],1)],1)],1)],1),l("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[l("el-button",{on:{click:function(e){return t.cancelSubmit("imeiForm")}}},[t._v("取 消")]),t.isDisabledbtn?t._e():l("el-button",{attrs:{type:"primary",disabled:t.disabledbtn},on:{click:function(e){t.submitForm("studentForm","post",t.Request.getVirtualDeviceUrl(),t.studentForm)}}},[t._v("确 定")])],1)],1),l("el-dialog",{attrs:{title:t.告警,visible:t.alarmVisible,width:"40%","close-on-click-modal":!1},on:{"update:visible":function(e){t.alarmVisible=e},closed:t.closedFormDialog}},[l("el-form",{directives:[{name:"loading",rawName:"v-loading",value:t.loadingDialog,expression:"loadingDialog"}],ref:"alarmForm",attrs:{model:t.alarmForm,rules:t.rules,"label-width":"120px"}},[l("el-row",{attrs:{gutter:20}},[l("el-col",{attrs:{span:20}},[l("el-form-item",{attrs:{label:"告警类型：",prop:"alarmType"}},[l("el-select",{attrs:{placeholder:"请选择告警类型"},model:{value:t.alarmForm.alarmType,callback:function(e){t.$set(t.alarmForm,"alarmType",e)},expression:"alarmForm.alarmType"}},t._l(t.selectAlarmType,(function(t){return l("el-option",{key:t.id,attrs:{label:t.alarmContent,value:t.id}})})),1)],1)],1)],1)],1),l("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[l("el-button",{on:{click:function(e){return t.cancelSubmit("imeiForm")}}},[t._v("取 消")]),t.isDisabledbtn?t._e():l("el-button",{attrs:{type:"primary",disabled:t.disabledbtn},on:{click:function(e){t.submitForm("alarmForm","post",t.Request.getAlarmUrl(),t.alarmForm)}}},[t._v("确 定")])],1)],1)],1)},a=[],o={name:"VirtualDevice",data:function(){return{Request:this.$api.api.prototype,Common:this.$common.common.prototype,loginVisible:!1,virtualDevices:[],selectSchool:[],dialogTitle:"",editVisible:!1,alarmVisible:!1,alarmType:"",studentForm:{studentNo:"",name:"",sex:"",age:"",phone:"",class:"",grade:"",schoolId:"",imei:""},selectAlarmType:[{id:1,alarmContent:"SOS告警"},{id:2,alarmContent:"低电量告警"},{id:3,alarmContent:"出围栏告警"},{id:4,alarmContent:"进围栏告警"}],alarmForm:{alarmType:"",imei:""}}},created:function(){this.getVirtualDevicesRequest(),this.getSchoolInfoDataRequest()},methods:{getVirtualDevicesRequest:function(){this.Common.getAxios(this.Request.getVirtualDeviceUrl(),"",this.getVirtualDevices)},getVirtualDevices:function(t){if(200===t.status){var e=t.data;this.virtualDevices=e.data,console.log("virtualDevices:",this.virtualDevices)}else 401===t.status?this.loginVisible=!0:this.$message.error("服务器繁忙，请稍后再试")},getSchoolInfoDataRequest:function(){this.Common.getAxios(this.Request.getSchoolInfoUrl(),"",this.getSchoolData)},getSchoolData:function(t){if(200===t.status){var e=t.data;this.loading=!1,this.selectSchool=e.data}else 401===t.status?this.loginVisible=!0:(this.$message.error("服务器繁忙，请稍后再试"),this.loading=!1)},addVirtualDevice:function(){this.dialogTitle="新增",this.editVisible=!0},submitForm:function(t,e,l,s){this.disabledbtn=!0,this.loadingDialog=!0,this.Common.commonAxios(l,e,s,this.submitFormInfo),this.disabledbtn=!1,this.loadingDialog=!1},submitFormInfo:function(t){if(200===t.status){var e=t.data;this.disabledbtn=!1,this.loadingDialog=!1,200===e.code?(this.editVisible=!1,this.alarmVisible=!1,this.$message.success("操作成功"),this.getVirtualDevicesRequest()):this.$message.error(e.message)}else 401===t.status?this.loginVisible=!0:(this.disabledbtn=!1,this.loadingDialog=!1,this.$message.error("服务器繁忙，请稍后再试"))},alarmRequest:function(t){this.alarmVisible=!0,this.alarmForm.imei=t}}},i=o,r=(l("9a15"),l("2877")),n=Object(r["a"])(i,s,a,!1,null,"009cb948",null);e["default"]=n.exports},"9a15":function(t,e,l){"use strict";l("ad17")},ad17:function(t,e,l){}}]);
//# sourceMappingURL=chunk-11c48044.874de31d.js.map