<template>
    <div>
        <van-nav-bar
                title="提交详情"
                left-text="返回"
                left-arrow
                @click-left="goBack"/>
        <van-row>
            <van-col span="22" offset="1">
                <van-panel :title="commit.taskTitle" :desc="commit.username" :status="commit.status"/>
            </van-col>
        </van-row>
        <van-row>
            <van-col span="22">
                <van-image
                        width="100"
                        height="100"
                        fit="contain"
                        :src="commit.commitImage"
                        @click="showImage([commit.commitImage])"/>
            </van-col>
        </van-row>
        <van-row>
            <van-field label="评论内容："
                       disabled
                       type="textarea"
                       autosize
                       v-model="commit.commentKey"/>
        </van-row>
        <van-row v-if="checkAdmin">
            <van-col offset="1" span="10">
                <van-button type="primary" :disabled="!hasStarted(commit.status)" round size="large"
                            @click="passCommit">通过
                </van-button>
            </van-col>
            <van-col offset="2" span="10">
                <van-button type="danger" :disabled="!hasStarted(commit.status)" round size="large"
                            @click="refuseCommit">拒绝
                </van-button>
            </van-col>
        </van-row>
    </div>
</template>
<script>
    import {ImagePreview} from 'vant'
    import {BACK_HOST, COMMITS} from '../../js/const/const'
    import {commitStatus} from "../../js/const/status";
    import {mapState} from 'vuex'
    import {ADMIN} from "../../js/const/admin";

    export default {
        data() {
            return {
                commit: {}
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
            this.commit = this.$route.params.commit
        },
        methods: {
            showImage: function(imgs) {
                ImagePreview({
                    images: imgs,
                });
            },
            goBack: function () {
                this.$router.go(-1)
            },
            hasStarted: function (status) {
                return status === commitStatus.commitStatusCommit;
            },
            passCommit: function () {
                this.updateCommitStatus(commitStatus.commitStatusDone)
            },
            refuseCommit: function () {
                this.updateCommitStatus(commitStatus.commitStatusRefuse)
            },
            updateCommitStatus: function (state) {
                let that = this;
                let commit = {
                    ID: that.commit.ID,
                    status: state
                };
                let url = BACK_HOST + that.user.name + "/" + that.commit.taskID + COMMITS + "?resolve=" + state;
                that.$axios.put(url, JSON.stringify(commit), {headers: {'Content-Type': 'application/json'}}).then((resp)=>{
                    console.log(resp.data);
                    that.goBack()
                }).catch((err)=>{
                    console.log(err)
                })
            }
        }
    }
</script>