<template>
    <div class="table">
        <div class="container">
            <el-collapse v-model="activeNames">
                <el-collapse-item title=""
                                  name="1">
                    <template slot="title">
                        <i class="header-icon el-icon-search"></i>查询条件：
                    </template>
                    <el-row :gutter="20">
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
                        <el-col :sm="5">
                            <el-select v-model="searchCondition.deviceCode"
                                       placeholder="设备类型"
                                       class="">
                                <el-option v-for="item in selectDeviceTypeCopy"
                                           :label="item.propListName"
                                           :value="item.propListCode"
                                           :key="item.id"></el-option>
                            </el-select>
                        </el-col>
                    </el-row>
                    <el-row :gutter="20">
                        <el-col :sm="5">
                            <el-input v-model="searchCondition.imei"
                                      placeholder="IMEI号"
                                      class=""></el-input>
                        </el-col>
                        <el-col :sm="5">
                            <el-input v-model="searchCondition.studentName"
                                      placeholder="学生姓名"
                                      class=""></el-input>
                        </el-col>
                        <el-col :sm="5">
                            <el-input v-model="searchCondition.phone"
                                      placeholder="手机号"
                                      class=""></el-input>
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
                </el-collapse-item>
                <el-collapse-item name="2">
                    <template slot="title">
                        <i class="header-icon el-icon-lx-cascades"></i>查询结果：
                    </template>
                    <div class="handle-box">
                        <el-button type="primary"
                                   icon="el-icon-folder-add"
                                   class="mr10"
                                   @click="addData">新增</el-button>
                        <el-button type="primary"
                                   icon="el-icon-folder-delete"
                                   class="mr10"
                                   @click="batchDelete">批量删除</el-button>
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
                                         align="center"></el-table-column>
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
                                <el-button type="text"
                                           icon="el-icon-edit"
                                           @click="handleViewEdit(scope.$index, scope.row,'edit','编辑')">编辑</el-button>
                                <el-button type="text"
                                           icon="el-icon-delete"
                                           class="red"
                                           @click="handleDelete(scope.$index, scope.row)">删除</el-button>
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
                </el-collapse-item>
            </el-collapse>
        </div>

        <!-- 新增编辑查看弹出框 -->
        <el-dialog :title="dialogTitle"
                   :visible.sync="editVisible"
                   width="40%"
                   :close-on-click-modal="false"
                   @closed="closedFormDialog">
            <el-form ref="imeiForm"
                     :model="studentForm"
                     :rules="rules"
                     v-loading="loadingDialog"
                     label-width="120px">
                <el-row :gutter="20">
                    <el-col :span="20">
                        <el-form-item label="学号："
                                      prop="studentNo">
                            <el-input v-model="studentForm.studentNo"
                                      placeholder="请输入学号"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="20">
                        <el-form-item label="姓名："
                                      prop="name">
                            <el-input v-model="studentForm.name"
                                      placeholder="请输入姓名"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="20">
                        <el-form-item label="年龄："
                                      prop="age">
                            <el-input v-model="studentForm.age"
                                      placeholder="请输入年龄"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-form-item label="性别："
                                  prop="sex">
                        <el-radio-group v-model="studentForm.sex">
                            <el-radio :label="1">男</el-radio>
                            <el-radio :label="0">女</el-radio>
                        </el-radio-group>
                    </el-form-item>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="20">
                        <el-form-item label="手机："
                                      prop="phone">
                            <el-input v-model="studentForm.phone"
                                      placeholder="请输入手机"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="20">
                        <el-form-item label="学校："
                                      prop="schoolId">
                            <el-select v-model="studentForm.schoolId"
                                       placeholder="请选择学校"
                                       class="">
                                <el-option v-for="item in selectSchool"
                                           :label="item.SchoolName"
                                           :value="item.ID"
                                           :key="item.ID"></el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="20">
                        <el-form-item label="年级："
                                      prop="grade">
                            <el-input v-model="studentForm.grade"
                                      placeholder="请输入年级"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="20">
                        <el-form-item label="班级："
                                      prop="class">
                            <el-input v-model="studentForm.class"
                                      placeholder="请输入班级"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="20">
                        <el-form-item label="IMEI号："
                                      prop="imei">
                            <el-input v-model="studentForm.imei"
                                      placeholder="请输入IMEI号"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-form>
            <span slot="footer"
                  class="dialog-footer">
        <el-button @click="cancelSubmit('imeiForm')">取 消</el-button>

        <el-button type="primary"
                   :disabled='disabledbtn'
                   v-if="!isDisabledbtn"
                   @click="submitForm('studentForm','post',studentUrl,studentForm)">确 定</el-button>
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
                   @click="deleteRow(deviceInfoUrl)">确 定</el-button>
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
        name: 'studentInfo',
        data () {
            return {
                Request: this.$api.api.prototype,/* 用于获取请求地址 */
                Common: this.$common.common.prototype,/* 用于获取公共方法 */
                studentUrl: '',
                gradeInfoUrl: '',
                classInfoUrl: '',
                deviceInfoUrl: '',
                schoolInfoUrl: '',
                deviceMoveInfoUrl: '',

                selectTableData: [],/* 选中行数据 */
                selectGradeCopy: [],
                selectClassCopy: [],
                selectSchoolCopy: [],
                selectSchool: [],
                selectSchoolTemplate: [],
                selectSchoolMove: [],
                selectSchoolMoveOld: [],
                selectSchoolMoveNew: [],
                selectDeviceTypeCopy: [],
                selectDeviceType: [],
                selectOperatorTypeCopy: [],
                selectOperatorType: [],

                activeNames: ['1', '2'],

                tableData: [],

                dialogTitle: '',
                editVisible: false,

                innerVisible: false,
                delVisible: false,
                loginVisible: false,/* 提示需要登录的对话框 */
                delVisibleType: 'delete',
                disabledbtn: false,/* 可用与不可用 */
                disabledbtnExport: false,/* 可用与不可用 */
                isDisabledbtn: false,/* 显示或隐藏 */

                importVisible: false,
                transferVisible: false,
                transferOldVisible: false,
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

                /* 新设备转移需要的数据 */
                transferForm: {
                    schoolId: '',
                    schoolName: ''
                },

                /* 旧设备转移需要的数据 */
                transferOldForm: {
                    srcId: '',
                    srcAccount: '',
                    destId: '',
                    destAccount: '',
                    imei: ''
                },

                searchCondition: {
                    schoolId: '',
                    gradeId: '',
                    classId: '',
                    studentName: '',
                    imei: '',
                    rfid: '',
                    phone: '',
                    deviceCode: '',
                    operatorType: '',
                    pageNumber: 1,/* 第几页 */
                    pageSize: 8/* 每页显示几条 */
                },

                studentForm: {
                    studentNo: '',
                    name: '',
                    sex: '',
                    age: '',
                    phone: '',
                    class: '',
                    grade: '',
                    schoolId: '',
                    imei: '',
                },
                rules: {
                    operatorType: [
                        { required: true, message: '请选择运营商', trigger: 'blur' }
                    ],
                    deviceCode: [
                        { required: true, message: '请选择设备类型', trigger: 'blur' }
                    ],
                    schoolId: [
                        { required: true, message: '请选择学校', trigger: 'blur' }
                    ],
                    imei: [
                        { required: true, message: '请输入IMEI号', trigger: 'blur' }
                    ],
                    rfid: [
                        { required: true, message: '请输入RFID号', trigger: 'blur' }
                    ]
                },
                /* 表头标题
                  参数，标题名称，是否选中，格式化方法
                */
                headerFilter: [
                    {
                        prop: "name",
                        label: "学生",
                        selected: true
                    },
                    {
                        prop: "student_no",
                        label: "学号",
                        selected: true
                    },
                    {
                        prop: "className",
                        label: "班级",
                        formatter: this.formatterClassName,
                        selected: true
                    },
                    {
                        prop: "school_name",
                        label: "学校",
                        selected: true
                    },
                    {
                        prop: "imei",
                        label: "IMEI号",
                        selected: true
                    },
                    {
                        prop: "phone",
                        label: "手机号",
                        selected: true
                    }
                ],
                idx: -1,
                id: -1
            }
        },
        created () {
            this.schoolInfoUrl = this.Request.getSchoolInfoUrl();
            this.studentUrl = this.Request.getStudentUrl();
            this.gradeInfoUrl = this.Request.getGradeInfoUrl();
            this.classInfoUrl = this.Request.getClassInfoUrl();
            /* 获取学校信息 */
            this.getSchoolInfoDataRequest();
            /* 获取表格信息 */
            this.getListDataRequest();
        },
        computed: {
        },

        methods: {
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
                return row.grade + "年级" + row.class + "班";
            },

            /* 请求：获取表格中数据 */
            getListDataRequest () {
                this.Common.getAxios(this.Request.getStudentInfoUrl(), '', this.getListData);
            },

            /* 请求：获取学校信息 */
            getSchoolInfoDataRequest () {
                this.Common.getAxios(this.Request.getSchoolInfoUrl(), '', this.getSchoolData);
            },

            /* 请求：获取设备类型信息 */
            getDeviceTypeRequest () {
                this.Common.getAxios(this.propListUrl, {
                    propTypeCode: 'DEVICE_MODEL'
                }, this.getPropList, ['selectDeviceTypeCopy', 'selectDeviceType']);
            },

            /* 请求：获取运营商类型信息 */
            getOperatorTypeRequest () {
                this.Common.getAxios(this.propListUrl, {
                    propTypeCode: 'OPERATOR_TYPE'
                }, this.getPropList, ['selectOperatorTypeCopy', 'selectOperatorType']);
            },

            /* 获取列表数据 */
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

            /* 获取学校数据  */
            getSchoolData (data) {
                if (data.status === 200) {
                    let tempData = data.data;
                    this.loading = false;
                    this.selectSchool = tempData.data;
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

            /* 查询区域：获取年级信息 */
            getGradeCopyInfoData (data) {
                if (data.status === 200) {
                    let tempData = data.data;
                    if (!this.validToken(tempData)) return;
                    if (tempData.code === 0 && tempData.result !== "") {
                        this.searchCondition.gradeId = '';
                        this.searchCondition.classId = '';
                        this.selectGradeCopy = [];
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
            /* 转移设备区域：学校下拉框事件 */
            changeSchoolTransfer (data) {
                let obj = {};
                /* 筛选出匹配数据 */
                obj = this.selectSchoolMove.find((item) => {
                    return item.id === data;
                });
                this.transferForm.schoolName = obj.schoolName;
            },
            /* 旧转移设备区域：学校下拉框事件 */
            changeSchoolTransferOld (data) {
                let obj = {};
                /* 筛选出匹配数据 */
                obj = this.selectSchoolMoveOld.find((item) => {
                    return item.id === data;
                });
                this.transferOldForm.srcAccount = obj.schoolName;
            },
            /* 旧移设备区域：学校下拉框事件 */
            changeSchoolTransferNew (data) {
                let obj = {};
                /* 筛选出匹配数据 */
                obj = this.selectSchoolMoveNew.find((item) => {
                    return item.id === data;
                });
                this.transferOldForm.destAccount = obj.schoolName;
            },

            /* 查询区域：年级下拉框事件 */
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
                this.studentForm.id = row.id;

                this.studentForm.deviceCode = row.deviceCode;
                this.studentForm.operatorType = row.operatorType;
                this.studentForm.schoolId = row.schoolId;
                this.studentForm.imei = row.imei;
                this.studentForm.rfid = row.rfid;
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
                this.searchCondition.id = '';
                this.searchCondition.operatorType = '';
                this.searchCondition.deviceCode = '';
                this.searchCondition.schoolId = '';
                this.searchCondition.gradeId = '';
                this.searchCondition.classId = '';
                this.searchCondition.studentName = '';
                this.searchCondition.imei = '';
                this.searchCondition.rfid = '';
                this.searchCondition.phone = '';
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

                this.studentForm.id = '';
                this.studentForm.deviceCode = '';
                this.studentForm.operatorType = '';
                this.studentForm.schoolId = '';
                this.studentForm.imei = '';
                this.studentForm.rfid = '';

                this.showType = 'add';
                this.methodType = 'post';
                this.$refs['imeiForm'].clearValidate();
            },

            /* 保存或编辑 */
            submitForm (formName, type, url, formData) {
                /* 防止多次提交 */
                this.disabledbtn = true;
                this.loadingDialog = true;
                this.Common.commonAxios(url, type, formData, this.submitFormInfo);
                this.disabledbtn = false;
                this.loadingDialog = false;
            },

            /* 保存或编辑信息 */
            submitFormInfo (data) {
                if (data.status === 200) {
                    let tempData = data.data;
                    this.disabledbtn = false;
                    this.loadingDialog = false;
                    if (tempData.code === 200) {
                        this.editVisible = false;
                        this.$message.success(tempData.message);
                        /* 重新刷新数据 */
                        this.getListDataRequest();
                    } else {
                        this.$message.error(tempData.message);
                    }
                } else if (data.status === 401) {
                    this.loginVisible = true;
                } else {
                    this.disabledbtn = false;
                    this.loadingDialog = false;
                    this.$message.error('服务器繁忙，请稍后再试');
                }
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
            },

            /* 关闭上传 弹出框事件 */
            closedImportDialog () {
                this.warningInfo = '';
                this.templateForm.schoolId = '';
                this.fileList = [];
            },

            /* 关闭新设备转移 弹出框事件 */
            closedTransferDialog () {
                this.transferForm.schoolId = '';
                this.transferForm.schoolName = '';
            },
            /* 关闭旧设备转移 弹出框事件 */
            closedTransferOldDialog () {
                this.transferOldForm.srcId = '';
                this.transferOldForm.srcAccount = '';
                this.transferOldForm.destId = '';
                this.transferOldForm.destAccount = '';
                this.transferOldForm.imei = '';
            },

            /* 取消提交 */
            cancelImport () {
                /* bug原因，不用api提供的方法 不信你试试 2.9.1版本未解决*/
                // this.$refs[formName].resetFields();
                this.importVisible = false;
            },

            /* 上传前判断文件的大小 */
            beforeUploadFile (file) {
                const isLt5M = file.size / 1024 / 1024 < 5;
                if (!isLt5M) {
                    this.$message.warning(`文件不能大于5M`);
                }
                return isLt5M;
            },

            /* 上传成功 */
            handleFileSuccess (data) {
                if (data.code !== 0) {
                    this.fileList = [];
                    this.$message.warning(data.message);
                    this.warningInfo = data.message;
                } else {
                    this.$message.success('批量导入成功');
                    this.getListDataRequest();
                }
            },

            /* 上传失败 */
            handleFileError () {
                this.$message.warning(`服务器繁忙，请稍后再试`);
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
    .header-icon {
        padding-right: 5px;
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

    /* 批量导入告警信息提示 */
    .warning-info {
        padding-top: 20px;
        color: #f00;
    }

    /* 提示图标设置 */
    .icon-wzg-config {
        height: 30px;
        line-height: 30px;
    }

    /* 筛选条件的布局 */
    .wzg-float-right {
        float: right;
    }
</style>
