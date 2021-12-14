<template>
    <div class="location">
        <el-col :span="6" class="side">
            <ul @mouseout="hoverAoi = null">
                <li
                        v-for="(item, i) in tracks"
                        :key="i"
                >
                    <el-card class="box-card">
                        <div slot="header" class="clearfix">
                            <span>{{item.studentName}}</span>
                            <el-button style="float: right;" type="primary" @click="track(item.track)">
                                轨迹回放
                            </el-button>
                        </div>
                      <div>
                        <ol style="margin-left: 10px;">
                          <li
                              v-for="(str, i) in item.trackStr"
                              :key="i"
                          >
                            {{str}}
                          </li>
                        </ol>
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
            >
                <amap-polyline
                    :path.sync="path"
                    :stroke-color="stroke"
                    :stroke-weight="7"
                    :strokeColor="'red'"
                    :cursor="'pointer'"
                >
                </amap-polyline>
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
        name: 'Track',
        data() {
            return {
                Request: this.$api.api.prototype,/* 用于获取请求地址 */
                Common: this.$common.common.prototype,/* 用于获取公共方法 */

                loginVisible: false,

                path: [],
                center: [114.4100667030411, 23.10351097195999],
                zoom: 18,

                tracks: [],

            };
        },

        created() {
            this.getTrackInfoRequest();
        },

        methods: {
            onHotspotClick(e) {
                if (e && e.lnglat) {
                    this.center = [e.lnglat.lng, e.lnglat.lat];
                }
            },

            getTrackInfoRequest() {
                this.Common.getAxios(this.Request.getTrackInfoUrl() , '', this.getTrackInfo);
            },

            getTrackInfo(data) {
                if (data.status === 200) {
                    let tempData = data.data;
                    this.tracks = tempData.data;
                } else if (data.status === 401) {
                    this.loginVisible = true;
                } else {
                    this.$message.error('服务器繁忙，请稍后再试');
                    this.loading = false;
                }
            },

            track(path) {
                console.log(path);
                this.path = path;
                if (path.length > 0) {
                    this.center = path[0];
                }
            },
        },
    };
</script>

<style scoped>
    .location {
        display: flex;
        height: 650px;
    }
    .map {
        height: 100%;
    }
    .side {
        height: 100%;
        overflow-y:scroll;
    }
</style>