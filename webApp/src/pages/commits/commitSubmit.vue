<template>
    <div>
        <van-nav-bar
                title="任务提交"
                left-text="返回"
                left-arrow
                @click-left="goBack"/>
        <br/>
        <van-col span="24">
            <van-cell title="请选择图片" :border="false"/>
        </van-col>
        <van-uploader
                v-model="imgData"
                :max-count="1"/>
        <br/>
        <van-button type="primary" v-on:click="onUpload">上传</van-button>
    </div>
</template>
<script>
    import {BACK_HOST, UPLOAD_IMAGES, COMMITS} from '../../js/const/const'
    import {mapState} from 'vuex'
    import {commitStatus} from "../../js/const/status";

    export default {
        data() {
            return {
                imgData: []
            }
        },
        computed: {
            ...mapState({
                user: state => state.user
            })
        },
        mounted() {
            console.log(this.$route.params.commit)
        },
        methods: {
            onUpload: function () {
                let that = this;
                if (that.imgData.length === 0) {
                    return that.$notify({
                        message: "请选择图片！",
                        type: "warning",
                    });
                }
                let formdata = new FormData();
                formdata.append("img0", that.imgData[0].file);
                that.$axios.post(BACK_HOST + that.user.name + UPLOAD_IMAGES, formdata, {headers: {'Content-Type': 'multipart/form-data'}}).then((resp)=>{
                    console.log(resp.data);
                    let commit = {
                        ID: that.$route.params.commit.ID,
                        commitImage: resp.data[0],
                        username: that.user.name,
                        taskID: this.$route.params.commit.taskID,
                        status: commitStatus.commitStatusCommit,
                    };
                    that.$axios.put(BACK_HOST + that.user.name + "/" + commit.taskID + COMMITS, JSON.stringify(commit), {headers: {'Content-Type': 'application/json'}}).then((resp)=>{
                        console.log(resp.data);
                        that.$notify({
                            message: "已提交，等待审核",
                            type: "success",
                        });
                        that.$router.push("/taskInfo");
                    }).catch((err)=>{
                        console.log(err)
                    })
                }).catch((err)=>{
                    console.log(err)
                })
            },
            goBack: function () {
                this.$router.go(-1)
            }
        }
    }
</script>