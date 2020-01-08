<template>
    <div>
        <van-nav-bar title="提交记录">
            <van-icon name="filter-o" slot="right" @click="showFilter=true"/>
        </van-nav-bar>
        <van-pull-refresh v-model="isLoading" @refresh="onRefresh">
            <van-list>
                <van-panel
                        v-for="c in commits"
                        :key="c.ID"
                        :title="c.taskTitle"
                        :desc="c.CreatedAt | transTime"
                        :status="c.status"
                        v-on:click="showCommitInfo(c)"/>
            </van-list>
        </van-pull-refresh>
        <van-popup
                v-model="showFilter"
                position="bottom"
                :style="{ height: '40%' }">
            <van-picker show-toolbar title="请选择"
                        :columns="commitStatus"
                        @cancel="showFilter=false"
                        @confirm="onConfirm"/>
        </van-popup>
    </div>
</template>
<script>
    import {BACK_HOST, COMMITS} from '../../js/const/const'
    import {mapState} from 'vuex'
    import {commitStatus} from "../../js/const/status";

    export default {
        data() {
            return {
                isLoading: false,
                finished: false,
                showFilter: false,
                commitStatus: ["全部", "进行中", "已提交", "已通过", "被拒绝"],
                filterVal: "全部",
                commits: []
            }
        },
        computed: {
            ...mapState({
                user: state => state.user
            })
        },
        mounted() {
            this.getData();
        },
        methods: {
            onRefresh: function () {
                let that = this;
                that.getData();
                that.isLoading = false
            },
            getData: function () {
                let that = this;
                let state = that.filterVal;
                let url = BACK_HOST + that.user.name + COMMITS;
                if ( state !== undefined && state !== "全部") {
                    url = url + "?state=" + state
                }
                that.$axios.get(url).then((resp) => {
                    that.commits = resp.data;
                }).catch((err) => {
                    console.log(err)
                })
            },
            onConfirm: function (val, index) {
                this.filterVal = val;
                this.getData();
                this.showFilter = false
            },
            showCommitInfo: function (commit) {
                console.log(commit);
                this.$router.push({
                    name: 'commitInfo',
                    params: {
                        commit: commit
                    },
                })
            }
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
        },
    }
</script>