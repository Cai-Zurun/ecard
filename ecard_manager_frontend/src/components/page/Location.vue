<template>
    <div class="location">
        <el-col :span="6" class="side">
            <ul @mouseout="hoverAoi = null" >
                <li
                        v-for="(location, i) in locations"
                        :key="i"
                >
                    <el-card class="box-card">
                        <div slot="header" class="clearfix">
                            <span>{{location.studentName}}</span>
                            <el-button type="primary" style="float: right;" @click="locate(location.longitude, location.latitude)">
                                定位
                            </el-button>
                        </div>
                        <div>
                            {{location.locationStr}}
                        </div>
                    </el-card>
                </li>
            </ul>
        </el-col>
        <el-col :span="18" class="map">
            <amap
                    cache-key="coord-picker-map"
                    mmap-style="amap://styles/whitesmoke"
                    async
                    :center.sync="center"
                    :zoom.sync="zoom"
                    is-hotspot
                    @hotspotclick="onHotspotClick"
            >
                <amap-marker :position="position"/>
            </amap>
        </el-col>


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
        name: 'Location',
        data() {
            return {
                Request: this.$api.api.prototype,/* 用于获取请求地址 */
                Common: this.$common.common.prototype,/* 用于获取公共方法 */

                loginVisible: false,

                center: [114.4100667030411, 23.10351097195999],
                position: [114.4100667030411, 23.10351097195999],   // [经度, 纬度]
                zoom: 18,

                locations: [],

            };
        },

        created() {
            this.getLocationInfoRequest();
        },

        methods: {
            onHotspotClick(e) {
                if (e && e.lnglat) {
                    this.center = [e.lnglat.lng, e.lnglat.lat];
                }
            },

            getLocationInfoRequest() {
                this.Common.getAxios(this.Request.getLocationInfoUrl() , '', this.getLocationInfo);
            },

            getLocationInfo(data) {
                if (data.status === 200) {
                    let tempData = data.data;
                    this.locations = tempData.data;
                } else if (data.status === 401) {
                    this.loginVisible = true;
                } else {
                    this.$message.error('服务器繁忙，请稍后再试');
                    this.loading = false;
                }
            },

            locate(longitude, latitude) {
                this.position = [longitude, latitude];
                this.center = [longitude, latitude];
            },

            load () {
              this.count += 2
            }
        },
    };
</script>

<style scoped>
    .location {
        display: flex;
        height: 600px;
    }
    .map {
        height: 100%;
    }
    .side {
        height: 100%;
        overflow-y:scroll;
    }
</style>