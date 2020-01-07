<template>
    <div>
        <van-nav-bar
                title="客服信息"
                left-text="返回"
                left-arrow
                @click-left="goBack">
            <van-icon v-if="checkAdmin" name="edit" slot="right" v-on:click="showUpdateService"/>
        </van-nav-bar>
        <van-row>
            <van-col span="24">
                <van-cell title="扫描二维码加客服微信" :border="false"/>
            </van-col>
        </van-row>
        <van-row>
            <van-col span="24">
                <van-image
                        width="30rem"
                        height="40rem"
                        fit="contain"
                        :src="serviceInfo.serviceQrCode"
                />
            </van-col>
        </van-row>
    </div>
</template>
<script>
    import {BACK_HOST, SERVICE_INFO} from "../../js/const/const";
    import {mapState} from 'vuex'
    import {ADMIN} from "../../js/const/admin";

    export default {
        data() {
            return {
                serviceInfo: {
                    serviceQrCode: "",
                    servicePhone: "",
                    serviceWechat: ""
                }
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
            let that = this;
            that.$axios.get(BACK_HOST + that.user.name + SERVICE_INFO).then((resp)=>{
                that.serviceInfo = resp.data
            }).catch((err)=>{
                console.log(err)
            })
        },
        methods: {
            goBack: function () {
                this.$router.go(-1)
            },
            showUpdateService: function () {
                this.$router.push({
                    name: 'updateService',
                    params: {
                        serviceInfo: this.serviceInfo
                    },
                })
            }
        }
    }
</script>