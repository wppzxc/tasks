<template>
    <div>register</div>
</template>

<script>
    import {USERS} from '../../js/const/const'
    import {mapActions} from 'vuex';
    export default {
        data() {
            return {
            }
        },
        mounted() {
            localStorage.clear();
            this.registerUser(this.$route.params.parent)
        },
        methods:{
            ...mapActions({
                loginFunc: 'loginFunc'
            }),
            registerUser: function (parent) {
                let that = this;
                let username = new Date().getTime();
                let url = USERS + username + "?parent=" + parent;
                that.$axios.post(url).then((resp)=>{
                    console.log(resp);
                    let user = resp.data;
                    that.loginFunc(user);
                    localStorage.setItem("username", user.name);
                    that.$router.push("/")
                }).catch((err)=>{
                    console.log(err);
                    that.$toast("注册失败：" + err)
                })
            }
        },
    }
</script>