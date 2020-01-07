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
        <van-row>
            <van-col offset="2" span="20">
                <van-button type="info" round size="large" v-on:click="onUpload">上传</van-button>
            </van-col>
        </van-row>
    </div>
</template>
<script>
    import {BACK_HOST, SERVICE_INFO, UPLOAD_IMAGES} from "../../js/const/const";
    import {mapState} from 'vuex'

    export default {
        data() {
            return {
                serviceInfo: {},
                imgData: []
            }
        },
        computed: {
            ...mapState({
                user: state => state.user
            })
        },
        mounted() {
            this.serviceInfo = this.$route.params.serviceInfo
        },
        methods: {
            afterRead: function (file) {
                this.imgData = [{
                    url: file.content,
                    isImage: true
                }];
            },
            onUpload: function () {
                let that = this;
                let formdata = new FormData();
                console.log(that.imgData);
                formdata.append("img0", that.dataURLtoFile(that.imgData[0].url, "img0.png"));
                that.$axios.post(BACK_HOST + that.user.name + UPLOAD_IMAGES, formdata, {headers: {'Content-Type': 'multipart/form-data'}}).then((resp)=>{
                    that.serviceInfo.serviceQrCode = resp.data[0];
                    that.$axios.put(BACK_HOST + that.user.name + SERVICE_INFO, that.serviceInfo).then((resp)=>{
                        that.goBack()
                    }).catch((err)=>{
                        console.log(err)
                    })
                })
            },
            goBack: function () {
                this.$router.go(-1)
            },
            dataURLtoFile: function(dataurl, filename) {//将base64转换为文件
                let arr = dataurl.split(','), mime = arr[0].match(/:(.*?);/)[1],
                    bstr = atob(arr[1]), n = bstr.length, u8arr = new Uint8Array(n);
                while(n--){
                    u8arr[n] = bstr.charCodeAt(n);
                }
                console.log([u8arr]);
                console.log(filename);
                return new File([u8arr], filename, {type:mime});
            }
        }
    }
</script>