<template>
    <div>
        <van-nav-bar
                title="任务提交"
                left-text="返回"
                left-arrow
                @click-left="goBack"/>
        <br/>
        <h3>请选择图片</h3>
        <van-uploader
                :after-read="afterRead"
                v-model="imgData"
                :max-count="1"/>
        <br/>
        <van-button type="primary" v-on:click="onUpload">上传</van-button>
    </div>
</template>
<script>
    export default {
        data() {
            return {
                imgData: []
            }
        },
        methods: {
            afterRead: function (file) {
                this.imgData = [{
                    url: file.content,
                    isImage: true
                }];
                console.log(this.imgData)
            },
            onUpload: function () {
                let that = this;
                setTimeout(function () {
                    that.$toast("上传成功！");
                    that.$router.go(-2)
                }, 1000)
            },
            goBack: function () {
                this.$router.go(-1)
            }
        }
    }
</script>