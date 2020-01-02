<template>
    <div>
        <van-nav-bar
                title="邀请信息"
                left-text="返回"
                left-arrow
                @click-left="goBack">
        </van-nav-bar>
        <van-cell title="扫描二维码" :border="false"/>
        <van-row>
            <van-col offset="7" span="10">
                <div id="qrcode" class="invite-qrcode"></div>
            </van-col>
        </van-row>
        <van-cell title="或" :border="false"/>
        <van-row>
            <van-col offset="2" span="20">
                <van-button type="primary" round size="large" v-on:click="copy">点击复制链接</van-button>
            </van-col>
        </van-row>
    </div>
</template>
<script>
    import  QRCode from 'qrcodejs2'
    import {INVITE_URL} from '../../js/const/const'
    import {mapState} from 'vuex'
    export default {
        data(){
            return {
                inviteURL: ""
            }
        },
        computed: {
            ...mapState({
                user: state => state.user
            })
        },
        mounted () {
            this.inviteURL = INVITE_URL + this.user.name;
            this.qrcode()
        },
        methods: {
            qrcode () {
                let that = this;
                let qrcode = new QRCode('qrcode', {
                    width: 100,
                    height: 100, // 高度
                    text: that.inviteURL // 二维码内容
                    // render: 'canvas' // 设置渲染方式（有两种方式 table和canvas，默认是canvas）
                    // background: '#f0f'
                    // foreground: '#ff0'
                });
                console.log(qrcode)
            },
            copy: function () {
                let that = this;
                let url = INVITE_URL + that.user.name;
                that.$copyText(url).then(function (e) {
                    that.$toast("已复制邀请链接，去邀请吧");
                }, function (e) {
                    that.$toast("无法获取邀请链接！");
                })
            },
            goBack: function () {
                this.$router.go(-1)
            },
        }
    }
</script>
<style>
    .invite-qrcode img {
        width: 100%;
    }
</style>