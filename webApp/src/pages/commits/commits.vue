<template>
    <div>
        <van-nav-bar title="完成记录">
            <van-icon name="filter-o" slot="right" @click="showFilter=true"/>
        </van-nav-bar>
        <van-pull-refresh v-model="isLoading" @refresh="onRefresh">
            <van-list
                    v-model="loading"
                    :finished="finished"
                    finished-text="没有更多了"
                    @load="onLoad"
            >
                <van-panel
                        v-for="c in commits"
                        :key="c.id"
                        :title="c.title"
                        :desc="c.description"
                        :status="c.status"
                        v-on:click="showCommitInfo(c)"/>
            </van-list>
        </van-pull-refresh>
        <van-popup
                v-model="showFilter"
                position="bottom"
                :style="{ height: '30%' }">
            <van-picker show-toolbar title="请选择"
                        :columns="commitStatus"
                        @cancel="showFilter=false"
                        @confirm="onConfirm" />
        </van-popup>
    </div>
</template>
<script>
    export default {
        data() {
            return {
                loading: false,
                isLoading: false,
                finished: false,
                showFilter: false,
                commitStatus: ["全部","进行中","通过","不通过"],
                commits: []
            }
        },
        beforeMount() {
            let cacheCommits = sessionStorage.getItem("commits");
            if (cacheCommits !== null) {
                this.commits = JSON.parse(cacheCommits);
            }
        },
        methods: {
            onRefresh: function () {
                let that = this;
                setTimeout(function () {
                    sessionStorage.setItem("commits", "");
                    let newCommits = [];
                    for (let i = 0; i < 20; i++) {
                        newCommits.push({
                            id: i,
                            title: "第 " + i + " 个任务",
                            description: "参与饿了么评论，领取奖励",
                            number: 100,
                            longContent: "你要这样这样这样做，再那样那样那样做，最后再这样这样这样做，就可以了",
                            status: "进行中",
                            images: [{
                                url: "https://img.yzcdn.cn/vant/cat.jpeg",
                            }, {
                                url: "https://img.yzcdn.cn/vant/t-thirt.jpg",
                            }, {
                                url: "https://img.yzcdn.cn/vant/cat.jpeg",
                            },{
                                url: "https://img.yzcdn.cn/vant/t-thirt.jpg",
                            }, {
                                url: "https://img.yzcdn.cn/vant/t-thirt.jpg",
                            }, {
                                url: "https://img.yzcdn.cn/vant/cat.jpeg",
                            }]
                        })
                    }
                    that.commits = newCommits;
                    sessionStorage.setItem("commits", JSON.stringify(newCommits));
                    that.$toast("success");
                    that.isLoading = false;
                    that.finished = false
                }, 1000);
            },
            onLoad: function () {
                let that = this;
                setTimeout(function () {
                    for (let i = 0; i < 20; i++) {
                        that.commits.push({
                            id: i,
                            title: "第 " + i + " 个任务",
                            description: "参与饿了么评论，领取奖励",
                            number: 100,
                            longContent: "你要这样这样这样做，再那样那样那样做，最后再这样这样这样做，就可以了",
                            status: "进行中",
                            images: [{
                                url: "https://img.yzcdn.cn/vant/cat.jpeg",
                            }, {
                                url: "https://img.yzcdn.cn/vant/t-thirt.jpg",
                            }, {
                                url: "https://img.yzcdn.cn/vant/cat.jpeg",
                            },{
                                url: "https://img.yzcdn.cn/vant/t-thirt.jpg",
                            }, {
                                url: "https://img.yzcdn.cn/vant/t-thirt.jpg",
                            }, {
                                url: "https://img.yzcdn.cn/vant/cat.jpeg",
                            }]
                        })
                    }
                    sessionStorage.setItem("commits", JSON.stringify(that.commits));
                    that.loading = false;
                    if (that.commits.length >= 100) {
                        that.finished = true
                    }
                }, 1000);
            },
            onConfirm: function(val, index) {
                this.$toast("已筛选：" + val);
                this.showFilter = false
            },
            showCommitInfo: function (commit) {
                this.$router.push({
                    name: 'commitInfo',
                    params: {
                        commit: commit
                    },
                })
            }
        }
    }
</script>