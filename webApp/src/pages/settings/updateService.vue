<template>
    <div>
        <van-nav-bar
                title="修改客服信息"
                left-text="返回"
                left-arrow
                @click-left="goBack">
            <van-icon name="edit" slot="right" />
        </van-nav-bar>
        <van-cell title="请选择二维码" :border="false"/>
        <van-uploader
                :after-read="afterRead"
                v-model="imgData"
                :max-count="1"/>
        <br/>
        <van-button type="info" round size="large" v-on:click="onUpload">上传</van-button>
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