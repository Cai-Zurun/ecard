<template>
    <div class="table">
        <div class="container">
            <div class="crumbs">
                <el-breadcrumb separator="/">
                    <el-breadcrumb-item><i class="el-icon-search"></i> 查询条件：</el-breadcrumb-item>
                </el-breadcrumb>
            </div>
            <el-row :gutter="20">
                <el-col :sm="5"
                        v-if="limitShow !== '1' && limitShow !=='2'">
                    <el-select v-model="searchCondition.conditionArray"
                               placeholder="告警类型"
                               class="">
                        <el-option v-for="item in selectAlarmType"
                                   :label="item.propListName"
                                   :value="item.propListCode"
                                   :key="item.id"></el-option>
                    </el-select>
                </el-col>
                <el-col :sm="5">
                    <el-select v-model="searchCondition.schoolId"
                               placeholder="请选择学校"
                               class=""
                               @change="changeSchoolCopy">
                        <el-option v-for="item in selectSchoolCopy"
                                   :label="item.schoolName"
                                   :value="item.id"
                                   :key="item.id"></el-option>
                    </el-select>
                </el-col>
                <el-col :sm="5">
                    <el-select v-model="searchCondition.gradeId"
                               placeholder="请选择年级"
                               class=""
                               @change="changeGradeCopy">
                        <el-option v-for="item in selectGradeCopy"
                                   :label="item.gradeName"
                                   :value="item.id"
                                   :key="item.id"></el-option>
                    </el-select>
                </el-col>
                <el-col :sm="5">
                    <el-select v-model="searchCondition.classId"
                               placeholder="请选择班级"
                               class="">
                        <el-option v-for="item in selectClassCopy"
                                   :label="item.className"
                                   :value="item.id"
                                   :key="item.id"></el-option>
                    </el-select>
                </el-col>
            </el-row>
            <el-row :gutter="20">
                <el-col :sm="5">
                    <el-input v-model="searchCondition.name"
                              placeholder="学生姓名"
                              class=""></el-input>
                </el-col>
                <el-col :sm="5">
                    <el-input v-model="searchCondition.phone"
                              placeholder="学生手机号"
                              class=""></el-input>
                </el-col>

                <el-col :sm="10">
                    <el-date-picker v-model="queryDate"
                                    type="datetimerange"
                                    range-separator="至"
                                    format="yyyy-MM-dd HH:mm:ss"
                                    value-format="yyyy-MM-dd HH:mm:ss"
                                    start-placeholder="开始日期"
                                    @change="changeQueryDate"
                                    end-placeholder="结束日期">
                    </el-date-picker>
                </el-col>
            </el-row>
            <el-row :gutter="20">
                <el-col :md="7">
                    <el-button type="primary"
                               icon="el-icon-search"
                               @click="searchInfo">搜索</el-button>
                    <el-button type="primary"
                               icon="el-icon-refresh"
                               @click="resetSearch">重置</el-button>
                </el-col>
            </el-row>
            <div class="crumbs">
                <el-breadcrumb separator="/">
                    <el-breadcrumb-item><i class="el-icon-lx-cascades"></i> 查询结果：</el-breadcrumb-item>
                </el-breadcrumb>
            </div>
            <div class="handle-box">
                <!-- 注意布局，此时没有其他按钮 el-popover组件用span包裹的，所以自己需要嵌套一个div进行布局-->
                <div style="height:24px">
                    <el-popover placement="bottom-end"
                                trigger="click">
                        <div v-for="(item,index) in headerFilter"
                             :key="index">
                            <el-checkbox v-model="item.selected">{{item.label}}</el-checkbox>
                        </div>
                        <el-button circle
                                   slot="reference"
                                   class="mr10 wzg-float-right"><i class="iconfont iconguolv"></i></el-button>
                    </el-popover>
                </div>
            </div>
            <el-table :data="tableData"
                      border
                      class="table"
                      ref="multipleTable"
                      v-loading="loading"
                      :reserve-selection="true"
                      @row-dblclick="rowDblclick">
                <el-table-column type="index"
                                 align="center">
                </el-table-column>
                <el-table-column type="selection"
                                 width="55"
                                 align="center">
                </el-table-column>
                <template v-for="(item, index) in headerFilter">
                    <el-table-column :prop="item.prop"
                                     :key="index"
                                     align="center"
                                     :label="item.label"
                                     :formatter="item.formatter"
                                     v-if="item.selected">
                    </el-table-column>
                </template>
                <el-table-column label="操作"
                                 width="180"
                                 align="center">
                    <template slot-scope="scope">
                        <el-button type="text"
                                   icon="el-icon-view"
                                   @click="handleViewEdit(scope.$index, scope.row,'view','查看')">查看</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="pagination">
                <el-pagination background
                               @size-change="handleSizeChange"
                               @current-change="handleCurrentChange"
                               layout="total, sizes, prev, pager, next, jumper"
                               :page-size="searchCondition.pageSize"
                               :current-page="searchCondition.pageNumber"
                               :page-sizes="pageSizes"
                               :total="total"
                               :hide-on-single-page="true">
                </el-pagination>
            </div>
        </div>

        <!-- 新增编辑查看弹出框 -->
        <el-dialog :title="dialogTitle"
                   :visible.sync="editVisible"
                   width="60%"
                   :close-on-click-modal="false"
                   @closed="closedFormDialog">
            <el-form ref="alarmForm"
                     :model="alarmInfoForm"
                     v-loading="loadingDialog"
                     label-width="100px">
                <el-row :gutter="20">
                    <el-col :span="10">
                        <el-form-item label="学校："
                                      prop="schoolName">
                            <el-input v-model="alarmInfoForm.schoolName"></el-input>
                        </el-form-item>
                    </el-col>
                    <el-col :span="10">
                        <el-form-item label="班级："
                                      prop="className">
                            <el-input v-model="alarmInfoForm.className"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>

                <el-row :gutter="20">
                    <el-col :span="10">
                        <el-form-item label="学生："
                                      prop="name">
                            <el-input v-model="alarmInfoForm.name"></el-input>
                        </el-form-item>
                    </el-col>
                    <el-col :span="10">
                        <el-form-item label="学生手机号："
                                      prop="phone">
                            <el-input v-model="alarmInfoForm.phone"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="10">
                        <el-form-item label="家长："
                                      prop="parentName">
                            <el-input v-model="alarmInfoForm.parentName"></el-input>
                        </el-form-item>
                    </el-col>
                    <el-col :span="10">
                        <el-form-item label="家长手机号："
                                      prop="parentPhone">
                            <el-input v-model="alarmInfoForm.parentPhone"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="10">
                        <el-form-item label="告警类型："
                                      prop="alarmName">
                            <el-input v-model="alarmInfoForm.alarmName"></el-input>
                        </el-form-item>
                    </el-col>
                    <el-col :span="10">
                        <el-form-item label="告警时间："
                                      prop="alarmTime">
                            <el-input v-model="alarmInfoForm.alarmTime"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="10">
                        <el-form-item label="经度："
                                      prop="lng">
                            <el-input v-model="alarmInfoForm.lng"></el-input>
                        </el-form-item>
                    </el-col>
                    <el-col :span="10">
                        <el-form-item label="纬度："
                                      prop="lat">
                            <el-input v-model="alarmInfoForm.lat"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="15">
                        <el-form-item label="地址："
                                      prop="content">
                            <el-input type="textarea"
                                      :rows="8"
                                      placeholder="地址"
                                      v-model="alarmInfoForm.address">
                            </el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-form>
            <span slot="footer"
                  class="dialog-footer">
        <el-button @click="cancelSubmit('alarmForm')">取 消</el-button>
        <el-button type="primary"
                   :disabled='disabledbtn'
                   v-if="!isDisabledbtn"
                   @click="submitForm('alarmForm',methodType,parentUrl,alarmInfoForm)">确 定</el-button>
      </span>
        </el-dialog>

        <!-- 删除提示框 -->
        <el-dialog title="提示"
                   :visible.sync="delVisible"
                   width="300px"
                   :close-on-click-modal='false'
                   center>
            <div class="del-dialog-cnt">删除不可恢复，是否确定删除？</div>
            <span slot="footer"
                  class="dialog-footer">
        <el-button @click="delVisible = false">取 消</el-button>
        <el-button type="primary"
                   :disabled='disabledbtn'
                   :class="{disabledbtn:isDisabledbtn}"
                   @click="deleteRow(parentUrl)">确 定</el-button>
      </span>
        </el-dialog>

        <!-- 重新登录提示框 -->
        <el-dialog title="提示"
                   :visible.sync="loginVisible"
                   width="300px"
                   :close-on-click-modal='false'
                   center>
            <div class="del-dialog-cnt">登录时间过长，请重新登录</div>
            <span slot="footer"
                  class="dialog-footer">
        <el-button @click="loginVisible = false">取 消</el-button>
        <el-button type="primary"
                   @click="Common.signOut('/login')">确 定</el-button>
      </span>
        </el-dialog>
    </div>
</template>

<script>
    export default {
        name: 'alarmInfo',
        data () {

            return {
                Request: this.$api.api.prototype,/* 用于获取请求地址 */
                Common: this.$common.common.prototype,/* 用于获取公共方法 */
                USERTYPE: 2,
                propListUrl: '',
                parentUrl: '',
                resetpwdInfoUrl: '',
                schoolInfoUrl: '',
                alarmInfoUrl: '',
                selectAlarmType: [],
                selectClassCopy: [],
                selectGradeCopy: [],
                selectClass: [],
                selectSchoolCopy: [],
                selectParentAccountCopy: [],
                selectSchool: [],
                selectSchoolTemplate: [],

                limitShow: localStorage.getItem('userType'),

                tableData: [],
                selectTableData: [],/* 选中行数据 */

                dialogTitle: '',
                editVisible: false,
                delVisible: false,
                loginVisible: false,/* 提示需要登录的对话框 */
                resetVisible: false,
                delVisibleType: 'delete',
                disabledbtn: false,/* 可用与不可用 */
                disabledbtnExport: false,/* 可用与不可用 */
                isDisabledbtn: false,/* 显示或隐藏 */

                importVisible: false,
                fileList: [],
                warningInfo: '',
                headers: {
                    Authorization: localStorage.getItem('token')
                },

                loading: true,
                loadingDialog: false,/* 对话框中数据提交加载 */
                propTypeCodeDisabled: false,/* 是否可编辑 */
                classCodeDisabled: true,/* 是否可编辑 */
                showType: 'add',/* 用于对话框显示 类型标记 */
                methodType: 'post',/* 请求类型 */

                total: 8,/* 总数 */
                /* 如果是后端分页 */
                //pageNumber: 1,/* 第几页 */
                //pageSize: 8,/* 每页显示几条 */
                pageSizes: [8, 20, 30, 40, 50, 100],

                /* 批量上传所需要的数据 */
                templateForm: {
                    schoolId: ''
                },

                queryDate: [],
                /* 查询条件 */
                searchCondition: {
                    conditionArray: '',
                    schoolId: '',
                    gradeId: '',
                    classId: '',
                    name: '',
                    phone: '',
                    parentName: '',
                    parentPhone: '',
                    startTime: '',
                    endTime: '',
                    pagination: 1,
                    pageNumber: 1,/* 第几页 */
                    pageSize: 8/* 每页显示几条 */
                },

                /* 表单中需要提交的数据 */
                alarmInfoForm: {
                    imei: '',
                    lat: '',
                    lng: '',
                    schoolName: '',
                    className: '',
                    alarmName: '',
                    alarmTime: '',
                    name: '',
                    phone: '',
                    parentName: '',
                    parentPhone: '',
                    address: ''
                },
                /* 表头标题
                 参数，标题名称，是否选中，格式化方法
               */
                headerFilter: [
                    {
                        prop: "alarmTime",
                        label: "告警时间",
                        // formatter: this.formatterAlarmTime,
                        selected: true
                    },
                    {
                        prop: "alarmType",
                        label: "告警类型",
                        selected: true
                    },
                    {
                        prop: "name",
                        label: "学生",
                        selected: true
                    },
                    {
                        prop: "schoolName",
                        label: "学校",
                        selected: true
                    },
                    {
                        prop: "class",
                        label: "班级",
                        formatter: this.formatterClassName,
                        selected: true
                    },
                    {
                        prop: "phone",
                        label: "学生手机号",
                        selected: true
                    },

                ],
                /* 删除时所需的编号 删除时的id */
                idx: -1,
                id: -1
            }
        },
        created () {
            this.alarmInfoUrl = this.Request.getAlarmInfoUrl();

            /* 获取表格信息 */
            this.getListDataRequest();
        },
        computed: {},
        methods: {
            /* 查询时间更改事件 */
            changeQueryDate (data) {
                if (data !== null) {
                    this.searchCondition.startTime = this.Common.renderTime(data[0]);
                    this.searchCondition.endTime = this.Common.renderTime(data[1]);
                } else {
                    this.searchCondition.startTime = '';
                    this.searchCondition.endTime = '';
                }
            },

            /* 每页显示几条 */
            handleCurrentChange (val) {
                this.searchCondition.pageNumber = val;
                this.getListDataRequest();
            },

            /* 显示第几页 */
            handleSizeChange (val) {
                this.searchCondition.pageSize = val;
                this.getListDataRequest();
            },

            /* 格式化班级 */
            formatterClassName (row) {
                return row.grade + '年级' + row.class + '班级';
            },

            /* 格式化告警时间 */
            formatterAlarmTime (row) {
                if (row.alarmTime !== null || row.alarmTime !== '') {
                    let temp = row.alarmTime.indexOf("-");
                    return row.alarmTime.substring(temp + 1);
                }
                return row.alarmTime;
            },

            /* 请求：获取公告类型 */
            getAlarmTypeRequest () {
                this.Common.getAxios(this.propListUrl, {
                    propTypeCode: 'ALARM_TYPE'
                }, this.getPropList, ['selectAlarmType']);
            },

            /* 请求：获取表格中数据 */
            getListDataRequest () {
                this.Common.getAxios(this.alarmInfoUrl , '', this.getListData);
            },

            /* 请求：获取学校信息 */
            getSchoolInfoDataRequest () {
                this.Common.getAxios(this.schoolInfoUrl, {
                    pagination: '0'
                }, this.getSchoolData);
            },

            /* 获取列表数据  */
            getListData (data) {
                if (data.status === 200) {
                    let tempData = data.data;
                    this.loading = false;
                    this.tableData = tempData.data;
                    this.total = tempData.data.length;
                } else if (data.status === 401) {
                    this.loginVisible = true;
                } else {
                    this.$message.error('服务器繁忙，请稍后再试');
                    this.loading = false;
                }
            },

            /* 获取属性列表数据 */
            getPropList (data, objArray) {
                if (data.status === 200) {
                    let tempData = data.data;
                    if (!this.validToken(tempData)) return;
                    if (tempData.code === 0 && tempData.result !== "") {
                        objArray.forEach(item => {
                            this[item] = tempData.result;
                        })
                    }
                } else {
                    this.$message.error('服务器繁忙，请稍后再试');
                }
            },

            /* 获取学校数据  */
            getSchoolData (data) {
                if (data.status === 200) {
                    let tempData = data.data;
                    this.loading = false;
                    if (!this.validToken(tempData)) return;
                    if (tempData.code === 0 && tempData.result !== "") {
                        this.selectSchoolCopy = tempData.result.rows;
                        this.selectSchool = tempData.result.rows;
                        this.selectSchoolTemplate = tempData.result.rows;
                    }
                } else {
                    this.$message.error('服务器繁忙，请稍后再试');
                    this.loading = false;
                }
            },
            /* 查询区域： 获取年级信息 */
            getGradeCopyInfoData (data) {
                if (data.status === 200) {
                    let tempData = data.data;
                    if (!this.validToken(tempData)) return;
                    if (tempData.code === 0 && tempData.result !== "") {
                        this.searchCondition.gradeId = '';
                        this.searchCondition.classId = '';
                        this.selectClassCopy = [];
                        this.selectGradeCopy = tempData.result;
                    }
                } else {
                    this.$message.error('服务器繁忙，请稍后再试');
                }
            },

            /* 查询区域：获取班级信息 */
            getClassInfoCopyData (data) {
                if (data.status === 200) {
                    let tempData = data.data;
                    if (!this.validToken(tempData)) return;
                    if (tempData.code === 0 && tempData.result !== "") {
                        this.searchCondition.classId = '';
                        this.selectClassCopy = tempData.result;
                    }
                } else {
                    this.$message.error('服务器繁忙，请稍后再试');
                }
            },

            /* 查询区域：学校下拉框事件 */
            changeSchoolCopy (data) {
                this.Common.getAxios(this.gradeInfoUrl, {
                    schoolId: data
                }, this.getGradeCopyInfoData);
            },

            /* 查询区域事件:  年级下拉框 */
            changeGradeCopy (data) {
                this.Common.getAxios(this.classInfoUrl, {
                    gradeId: data
                }, this.getClassInfoCopyData);
            },

            /* 双击选中当前行 */
            rowDblclick (row) {
                this.$refs.multipleTable.toggleRowSelection(row);
            },

            /* 点击查看或编辑按钮 */
            handleViewEdit (index, row, code, name) {
                if ('view' === code) {
                    this.dialogTitle = name;
                    this.isDisabledbtn = true;
                } else if ('edit' === code) {
                    this.dialogTitle = name;
                    this.methodType = 'put';
                }
                this.alarmInfoForm.id = row.id;

                this.alarmInfoForm.imei = row.imei;
                this.alarmInfoForm.lat = row.lat;
                this.alarmInfoForm.lng = row.lng;
                this.alarmInfoForm.schoolName = row.school.schoolName;
                this.alarmInfoForm.className = row.grade.gradeName + row.class.className;
                this.alarmInfoForm.alarmName = row.alarmName;
                this.alarmInfoForm.alarmTime = row.alarmTime;
                this.alarmInfoForm.name = row.name;
                this.alarmInfoForm.phone = row.phone;
                this.alarmInfoForm.parentName = row.parent.realName;
                this.alarmInfoForm.parentPhone = row.parent.phone;
                this.alarmInfoForm.address = row.address;

                this.showType = code;
                this.propTypeCodeDisabled = true;
                this.classCodeDisabled = true;
                this.editVisible = true;
            },

            /* 点击删除按钮 */
            handleDelete (index, row) {
                this.idx = index;
                this.id = row.id;
                this.delVisible = true;
                this.delVisibleType = 'delete';
            },

            /* 搜索 */
            searchInfo () {
                /* 重置到第一页 */
                this.loading = true;
                this.searchCondition.pageNumber = 1;
                this.getListDataRequest();
            },

            /* 重置搜索内容 */
            resetSearch () {
                this.searchCondition.conditionArray = '';
                this.searchCondition.schoolId = '';
                this.searchCondition.gradeId = '';
                this.searchCondition.classId = '';
                this.searchCondition.name = '';
                this.searchCondition.phone = '';
                this.searchCondition.parentName = '';
                this.searchCondition.parentPhone = '';
                this.searchCondition.startTime = '';
                this.searchCondition.endTime = '';
                this.queryDate = [];
                this.searchInfo();
            },

            /* 新增 */
            addData () {
                this.dialogTitle = '新增';
                this.editVisible = true;
            },

            /* 批量删除 */
            batchDelete () {
                /* 获取选中的数据 */
                let selection = this.$refs.multipleTable.store.states.selection;
                this.selectTableData = selection;
                if (selection.length > 0) {
                    this.delVisible = true;
                    this.delVisibleType = 'batchDelete';
                } else {
                    this.$message.warning('请至少选择一条数据');
                }
            },

            /* 关闭新增编辑查看 弹出框事件 */
            closedFormDialog () {
                /* 初始化数据使用clearValidate清空校验，
                  resetFields方法在点击编辑而后点击新增时是无效的，
                  所以还是要手动初始化所有的值
                */
                this.isDisabledbtn = false;
                this.propTypeCodeDisabled = false;
                this.classCodeDisabled = true;

                this.alarmInfoForm.id = '';
                this.alarmInfoForm.schoolId = '';
                this.alarmInfoForm.parentName = '';
                this.alarmInfoForm.phone = '';
                this.alarmInfoForm.intelligent = '';
                this.showType = 'add';
                this.methodType = 'post';
                this.$refs['alarmForm'].clearValidate();
            },

            /* 保存或编辑 */
            submitForm (formName, type, url, formData) {
                /* 防止多次提交 */
                this.disabledbtn = true;
                this.loadingDialog = true;
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        this.Common.commonAxios(url, type, formData, this.submitFormInfo);
                    } else {
                        this.disabledbtn = false;
                        this.loadingDialog = false;
                        return false;
                    }
                });
            },

            /* 保存或编辑信息 */
            submitFormInfo (data) {
                if (data.status === 200) {
                    let tempData = data.data;
                    this.disabledbtn = false;
                    this.loadingDialog = false;
                    if (!this.validToken(tempData)) return;
                    if (tempData.code === 0) {
                        this.editVisible = false;
                        this.$message.success(tempData.message);
                        /* 重新刷新数据 */
                        this.getListDataRequest();
                    } else {
                        this.$message.error(tempData.message);
                    }
                } else {
                    this.disabledbtn = false;
                    this.loadingDialog = false;
                    this.$message.error('服务器繁忙，请稍后再试');
                }
            },

            /* 保存或编辑 取消提交 */
            cancelSubmit () {
                /* bug原因，不用api提供的方法 不信你试试 2.9.1版本未解决*/
                // this.$refs[formName].resetFields();
                this.editVisible = false;
            },

            /* 确定删除 */
            deleteRow (deleteUrl) {
                if (this.delVisibleType !== 'delete' && this.delVisibleType !== 'batchDelete') {
                    this.delVisible = false;
                    this.$message.error(`别偷偷的修改特定的参数哦`);
                    return;
                }
                let selection = [];
                let length = 0;
                /* 获取要删除数据的id值数组 (单个与批量删除相同) */
                let idArray = [];
                if (this.delVisibleType === 'delete') {
                    idArray.push(this.id);
                } else if (this.delVisibleType === 'batchDelete') {
                    /*获取选中的值 */
                    selection = this.$refs.multipleTable.store.states.selection;
                    length = selection.length;
                    for (let i = 0; i < length; i++) {
                        idArray.push(selection[i].id);
                    }
                }
                this.loading = true;
                this.disabledbtn = true;
                /* 删除一条或多条数据 */
                this.Common.deleteAxios(deleteUrl, {
                    idArray: idArray
                }, this.deleteInfo, idArray);
                this.delVisible = false;
            },

            /* 删除后的信息 */
            deleteInfo (data, idArray) {
                this.disabledbtn = false;
                this.loading = false;
                if (data.status === 200) {
                    let tempData = data.data;
                    if (!this.validToken(tempData)) return;
                    if (tempData.code !== 0) {
                        this.$message.error(tempData.message);
                        return;
                    }
                    /* 当删除成功后，删除对应选中的数据或单个数据 */
                    for (let j = this.tableData.length - 1; j >= 0; j--) {
                        for (let i = 0; i < idArray.length; i++) {
                            if (idArray[i] === this.tableData[j].id) {
                                this.tableData.splice(j, 1);
                                break;
                            }
                        }
                    }
                    this.$message.success(`成功删除了${idArray.length}条数据`);
                    this.selectTableData = [];/* 清除选中的数据 */
                } else {
                    this.$message.error('服务器繁忙，请稍后再试');
                }
            },

            /* 校验token */
            validToken (data) {
                if (data.code === 2001 || data.code === 2101 || data.code === 2102) {
                    this.loginVisible = true;
                    return false;
                }
                return true;
            }
        }
    }

</script>

<style scoped>
    /* layout-布局 */
    .el-row {
        margin-bottom: 20px;
    }
    .el-row:last-child {
        margin-bottom: 0;
    }
    .el-col {
        border-radius: 4px;
    }
    .bg-purple-dark {
        background: #99a9bf;
    }
    .bg-purple {
        background: #d3dce6;
    }
    .bg-purple-light {
        background: #e5e9f2;
    }
    .grid-content {
        border-radius: 4px;
        min-height: 36px;
    }
    .row-bg {
        padding: 10px 0;
        background-color: #f9fafc;
    }

    .handle-box {
        margin-bottom: 20px;
    }

    .handle-select {
        width: 120px;
    }

    .handle-input {
        width: 300px;
        display: inline-block;
    }
    .del-dialog-cnt {
        font-size: 16px;
        text-align: center;
    }
    .table {
        width: 100%;
        font-size: 14px;
    }
    .red {
        color: #ff0000;
    }
    .mr10 {
        margin-right: 10px;
    }

    .textAlignCenter {
        text-align: center;
    }

    /* 修改select的盒模型 */
    .el-select {
        display: block;
    }

    /* 确定按钮的显示与隐藏 */
    .disabledbtn {
        display: none;
    }

    /* 修改该属性值  查询条件，查询结果得样式 */
    .el-breadcrumb {
        border-bottom: 1px solid #ebeef5;
        padding-bottom: 10px;
    }

    /* 提示图标设置 */
    .icon-wzg-config {
        height: 30px;
        line-height: 30px;
    }

    /* 批量导入告警信息提示 */
    .warning-info {
        padding-top: 20px;
        color: #f00;
    }

    /* 筛选条件的布局 */
    .wzg-float-right {
        float: right;
    }
</style>
