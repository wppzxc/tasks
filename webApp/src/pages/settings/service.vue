<template>
    <div>
        <van-nav-bar
                title="客服信息"
                left-text="返回"
                left-arrow
                @click-left="goBack">
            <van-icon name="edit" slot="right" v-on:click="showUpdateService"/>
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
                        :src="serviceImageUrl"
                />
            </van-col>
        </van-row>
    </div>
</template>
<script>
    export default {
        data() {
            return {
                serviceImageUrl: "https://gss0.bdstatic.com/94o3dSag_xI4khGkpoWK1HF6hhy/baike/c0%3Dbaike80%2C5%2C5%2C80%2C26/sign=fa9140accd95d143ce7bec711299e967/2934349b033b5bb571dc8c5133d3d539b600bc12.jpg"
            }
        },
        methods: {
            goBack: function () {
                this.$router.go(-1)
            },
            showUpdateService: function () {
                this.$router.push("/updateService")
            }
        }
    }
</script>