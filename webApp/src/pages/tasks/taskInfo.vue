<template>
    <div>
        <van-nav-bar title="任务详情">
            <van-icon v-if="checkAdmin" name="edit" slot="right" v-on:click="showEditTask"/>
        </van-nav-bar>
        <van-panel :title="taskInfo.title" :desc="taskInfo.expire | transTime" :status="taskInfo.status">
            <van-cell title="任务赏金：￥" :value="taskInfo.bonus | transMoney"/>
            <div>
                <van-cell title="任务须知"/>
                <van-field
                        :border="false"
                        type="textarea"
                        disabled
                        autosize
                        v-for="(s, key) in tmpSteps"
                        :key="key"
                        :value="s"/>
            </div>
            <div>
                <van-cell title="详细步骤"/>
                <van-grid :column-num="3">
                    <van-grid-item v-for="(image, key) in tmpGuideImages" :key="key">
                        <van-image :src="image" @click="showImages(tmpGuideImages, key)"/>
                    </van-grid-item>
                </van-grid>
            </div>
            <div>
                <van-cell title="任务内容"/>
                <van-grid :column-num="3">
                    <van-grid-item v-for="(image, key) in tmpCommentImages" :key="key">
                        <van-image :src="image" @click="showImages(tmpCommentImages, key)"/>
                    </van-grid-item>
                </van-grid>
                <van-field label="评论内容："
                           placeholder="请开始任务，获取评论内容"
                           disabled
                           type="textarea"
                           autosize
                           v-model="currentCommit.commentKey"/>
            </div>
            <van-row>
                <van-col offset="1" span="10">
                    <van-button type="primary" :disabled="hasStarted(currentCommit.status)" round size="large"
                                @click="startTask">开始
                    </van-button>
                </van-col>
                <van-col offset="2" span="10">
                    <van-button type="primary" :disabled="!hasStarted(currentCommit.status)" round size="large"
                                @click="goSubmitCommit">完成
                    </van-button>
                </van-col>
            </van-row>
        </van-panel>
    </div>
</template>
<script>
    import {ImagePreview} from 'vant'
    import {BACK_HOST, COMMITS, TASKS} from '../../js/const/const'
    import {commitStatus} from '../../js/const/status'
    import {mapState} from 'vuex'
    import {ADMIN} from "../../js/const/admin";
    export default {
        data() {
            return {
                taskInfo: {},
                tmpGuideImages: [],
                tmpCommentImages: [],
                tmpSteps: [],
                currentCommit: {
                    ID: 0,
                    commentKey: "",
                    commitImage: "",
                    status: "",
                },
            }
        },
        computed: {
            ...mapState({
                user: state => state.user
            }),
            checkAdmin: function () {
                return this.user.name === ADMIN
            }
        },
        mounted() {
            this.task = this.$route.params.task;
            let that = this;
            let username = localStorage.getItem("username");
            that.$axios.get(BACK_HOST + username + TASKS).then((resp) => {
                that.taskInfo = resp.data;
                that.tmpGuideImages = JSON.parse(that.taskInfo.guideImages);
                that.tmpCommentImages = JSON.parse(that.taskInfo.commentImages);
                that.tmpSteps = JSON.parse(that.taskInfo.steps);
                that.$axios.get(BACK_HOST + username + "/" + that.taskInfo.ID + COMMITS).then((resp) => {
                    let currentCommit = resp.data;
                    if (currentCommit.status === commitStatus.commitStatusStart) {
                        that.currentCommit = currentCommit
                    }
                }).catch((err) => {
                    console.log(err)
                })
            }).catch((err) => {
                console.log(err)
            })
        },
        methods: {
            showImages: function (images, index) {
                ImagePreview({
                    images: images,
                    startPosition: index,
                });
            },
            goBack: function () {
                this.$router.go(-1)
            },
            showEditTask: function () {
                this.$router.push('/editTask')
            },
            startTask: function () {
                let that = this;
                if (!that.user.registered) {
                    return that.$notify({
                        message: "请先绑定支付宝！",
                        type: "warning",
                    });
                }
                let username = localStorage.getItem("username");
                that.$axios.post(BACK_HOST + username + "/" + that.taskInfo.ID + COMMITS).then((resp) => {
                    console.log(resp.data);
                    that.currentCommit = resp.data;
                    that.$notify({
                        message: "已开始，等待完成",
                        type: "success",
                    });
                }).catch((err) => {
                    console.log(err)
                })
            },
            hasStarted: function (status) {
                return status === commitStatus.commitStatusStart;
            },
            goSubmitCommit: function () {
                let that = this;
                that.$router.push({
                    name: 'commitSubmit',
                    params: {
                        commit: that.currentCommit
                    },
                })
            },
        },
        filters: {
            transTime: function (timestamp) {
                let date = new Date(timestamp);//时间戳为10位需*1000，时间戳为13位的话不需乘1000
                let Y = date.getFullYear() + '-';
                let M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-';
                let D = date.getDate() + ' ';
                let h = date.getHours() + ':';
                let m = date.getMinutes() + ':';
                let s = date.getSeconds();
                return Y + M + D + h + m + s;
            },
            transMoney: function (money) {
                return (money/100).toFixed(2)
            }
        },
    }
</script>
<style>
    .task-info {
        font-size: xx-large;
    }
</style>