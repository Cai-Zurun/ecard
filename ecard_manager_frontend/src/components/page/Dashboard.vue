<template>
    <div class="index-img">
        <el-row :gutter="20">
            <el-col :span="8">
                <el-card shadow="hover"
                         class="mgb20"
                         style="height:252px;">
                    <div class="lc2-details">设备状态统计</div>
                    <g2-pie :type="'ring'" :axis-name="{name:'类别', value:'人次(次)'}"
                            :data="[{name:'在线', value:this.onlineCount},{name:'离线', value:this.offlineCount}]"
                            :guide="{name:'在线', value:this.onlinePercentage}"
                            class="pie-chart">
                    </g2-pie>
                </el-card>
                <el-card shadow="hover"
                         style="height:312px;">
                    <div slot="header"
                         class="clearfix">
                        <span>告警排行</span>
                    </div>
                    <el-col :span="24">
                      <g2-column :is-bar="false"
                                 :data="alarmRank"
                                 :axis-name="{name:'告警', value:'个数'}"
                                 :single-color="'#1890ff'"
                                 class="bar-chart" >
                      </g2-column>
                    </el-col>
                </el-card>
            </el-col>
            <el-col :span="16">
                <el-row :gutter="20"
                        class="mgb20">
                    <el-col :span="8">
                            <el-card shadow="hover"
                                     :body-style="{padding: '0'}">
                                <div class="grid-content grid-con-1">
                                    <i class="el-icon-s-custom grid-con-icon"></i>
                                    <div class="grid-cont-right">
                                        <div class="grid-num">{{studentCount}}</div>
                                        学生总数
                                    </div>
                                </div>
                            </el-card>
                    </el-col>
                    <el-col :span="8">
                            <el-card shadow="hover"
                                     :body-style="{padding: '0'}">
                                <div class="grid-content grid-con-2">
                                    <i class="el-icon-mobile grid-con-icon"></i>
                                    <div class="grid-cont-right">
                                        <div class="grid-num">{{deviceCount}}</div>
                                        设备总数
                                    </div>
                                </div>
                            </el-card>
                    </el-col>
                    <el-col :span="8">
                            <el-card shadow="hover"
                                     :body-style="{padding: '0'}">
                                <div class="grid-content grid-con-3">
                                    <i class="el-icon-message-solid grid-con-icon"></i>
                                    <div class="grid-cont-right">
                                        <div class="grid-num">{{alarmCount}}</div>
                                        告警总数
                                    </div>
                                </div>
                            </el-card>
                    </el-col>
                </el-row>
                <el-card shadow="hover"
                         style="height:463px;">
                    <div slot="header"
                         class="clearfix">
                        <span>设备分布</span>
                    </div>
                    <div style="height:463px;">
                        <amap
                                cache-key="coord-picker-map"
                                mmap-style="amap://styles/whitesmoke"
                                async
                                :center.sync="center"
                                :zoom.sync="zoom"
                                is-hotspot
                                @hotspotclick="onHotspotClick"
                        >
                            <div
                                    v-for="(postion, i) in positions"
                                    :key="i"
                            >
                                <amap-marker  :position="postion"/>
                            </div>
                        </amap>
                    </div>
                </el-card>
            </el-col>
        </el-row>

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
        name: 'dashboard',
        data () {
            return {
                Request: this.$api.api.prototype,/* 用于获取请求地址 */
                Common: this.$common.common.prototype,/* 用于获取公共方法 */
                loginVisible: false,/* 提示需要登录的对话框 */
                studentCount: '',
                deviceCount: '',
                onlineCount: 1,
                offlineCount: 0,
                onlinePercentage: '100%',
                alarmCount: '',


                loading: true,
                total: 0,/* 总数 */
                tableData: [],
                /* 查询条件 */
                searchCondition: {
                    pageNumber: 1,/* 第几页 */
                    pageSize: 4/* 每页显示几条 */
                },

                alarmRank:[],

                loading1: true,
                total1: 0,/* 总数 */
                tableData1: [],
                /* 查询条件 */
                searchCondition1: {
                    attendType: 1,
                    statusType: 1,
                    nowTime: '',
                    pageNumber: 1,/* 第几页 */
                    pageSize: 4/* 每页显示几条 */
                },

                center: [114.4100667030411, 23.10351097195999],
                positions: [],
                zoom: 18,
            }
        },
        computed: {
        },
        created () {
            this.getStudentsCountRequest();
            this.getDevicesCountRequest();
            this.getAlarmsCountRequest();
            this.getLocationInfoRequest();
            this.getAlarmRankRequest();
        },
        mounted () {
        },
        activated () {
        },

        methods: {
            // 获取学生人数
            getStudentsCountRequest () {
                this.Common.getAxios(this.Request.getStudentsCountUrl(), '', this.getStudentsCount);
            },
            getStudentsCount (data) {
                if (data.status === 200) {
                    let tempData = data.data;
                    if (tempData.code === 200 && tempData.data !== "") {
                        this.studentCount = tempData.data;
                    }
                } else if (data.status === 401) {
                    this.loginVisible = true;
                } else {
                    this.$message.error('服务器繁忙，请稍后再试');
                }
            },

            // 获取设备总数
            getDevicesCountRequest () {
                this.Common.getAxios(this.Request.getDevicesCountUrl(), '', this.getDevicesCount);
            },
            getDevicesCount (data) {
                if (data.status === 200) {
                    let tempData = data.data;
                    if (tempData.code === 200 && tempData.data !== "") {
                        this.deviceCount = tempData.data.deviceCount;
                        this.onlineCount = parseInt(tempData.data.onlineCount);
                        this.offlineCount = parseInt(tempData.data.offlineCount);
                        this.onlinePercentage = Math.round(this.onlineCount/(this.onlineCount+this.offlineCount)*100) +'%'
                    }
                } else if (data.status === 401) {
                    this.loginVisible = true;
                } else {
                    this.$message.error('服务器繁忙，请稍后再试');
                }
            },

            // 获取告警总数
            getAlarmsCountRequest () {
                this.Common.getAxios(this.Request.getAlarmsCountUrl(), '', this.getAlarmsCount);
            },
            getAlarmsCount (data) {
                if (data.status === 200) {
                    let tempData = data.data;
                    if (tempData.code === 200 && tempData.data !== "") {
                        this.alarmCount = tempData.data;
                    }
                } else if (data.status === 401) {
                    this.loginVisible = true;
                } else {
                    this.$message.error('服务器繁忙，请稍后再试');
                }
            },

            // 获取设备定位
            getLocationInfoRequest() {
                this.Common.getAxios(this.Request.getLocationInfoUrl() , '', this.getLocationInfo);
            },
            getLocationInfo(data) {
                if (data.status === 200) {
                    let locations = data.data.data;
                    for (let i = 0; i < locations.length; i++) {
                        this.positions.push([locations[i].longitude, locations[i].latitude])
                    }
                } else if (data.status === 401) {
                    this.loginVisible = true;
                } else {
                    this.$message.error('服务器繁忙，请稍后再试');
                    this.loading = false;
                }
            },

            // 获取设备定位
            getAlarmRankRequest() {
              this.Common.getAxios(this.Request.getAlarmRankUrl() , '', this.getAlarmRank);
            },
            getAlarmRank(data) {
              if (data.status === 200) {
                this.alarmRank = data.data.data;
              } else if (data.status === 401) {
                this.loginVisible = true;
              } else {
                this.$message.error('服务器繁忙，请稍后再试');
                this.loading = false;
              }
            },
        }
    }
</script>
<style scoped>
    .el-row {
        margin-bottom: 20px;
    }

    .grid-content {
        display: flex;
        align-items: center;
        height: 100px;
    }
    .teacher-h {
        height: 130px;
    }
    .grid-cont-right {
        flex: 1;
        text-align: center;
        font-size: 14px;
        color: #999;
    }

    .grid-num {
        font-size: 30px;
        font-weight: bold;
    }

    .grid-con-icon {
        font-size: 50px;
        width: 100px;
        height: 100px;
        text-align: center;
        line-height: 100px;
        color: #fff;
    }

    .grid-con-1 .grid-con-icon {
        background: rgb(45, 140, 240);
    }

    .grid-con-1 .grid-num {
        color: rgb(45, 140, 240);
    }

    .grid-con-2 .grid-con-icon {
        background: rgb(100, 213, 114);
    }

    .grid-con-2 .grid-num {
        color: rgb(100, 213, 114);
    }

    .grid-con-3 .grid-con-icon {
        background: rgb(242, 94, 67);
    }

    .grid-con-3 .grid-num {
        color: rgb(242, 94, 67);
    }

    .user-info {
        display: flex;
        align-items: center;
        padding-bottom: 20px;
        border-bottom: 2px solid #ccc;
        margin-bottom: 20px;
    }

    .user-avator {
        width: 120px;
        height: 120px;
        border-radius: 50%;
        object-fit: cover;
    }

    .user-info-cont {
        padding-left: 30px;
        flex: 1;
        font-size: 14px;
        color: #999;
    }
    .user-info-cont div:first-child {
        font-size: 30px;
        color: #222;
    }
    .user-info-list {
        font-size: 14px;
        color: #999;
        line-height: 25px;
    }

    .user-info-list span {
        margin-left: 10px;
    }

    .mgb20 {
        margin-bottom: 20px;
    }

    .todo-item {
        font-size: 14px;
    }

    .todo-item-del {
        text-decoration: line-through;
        color: #999;
    }

    .schart {
        width: 100%;
        height: 300px;
    }

    .index-img {
        width: 100%;
        /* height: 100%; */
    }
    .img-test {
        width: 100%;
        height: 100%;
        /* margin: auto; */
    }
    .user-info-name {
        margin-bottom: 10px;
    }

    .course-class {
        display: inline-block;
        /* float: left; */
        margin: 5px;
        padding: 5px;
        background: rgb(100, 213, 114);
        color: #fff;
        /* border: 1px solid red; */
        border-radius: 5px;
    }
    .course-class-no {
        background: #666;
    }
    .handle-box {
        margin-bottom: 20px;
    }
    .pie-chart {
        height: 200px;
    }
    .lc2-details {
        font-size: 22px;
        font-weight: bold;
        display: flex;
        justify-content:center;
    }
    .bar-chart {
        height: 200px;
    }
</style>
