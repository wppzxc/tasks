<template>
    <div>
        <van-nav-bar title="任务列表">
            <van-icon name="plus" slot="right" v-on:click="showAddTask"/>
        </van-nav-bar>
        <van-pull-refresh v-model="isLoading" @refresh="onRefresh">
            <van-list
                    v-model="loading"
                    :finished="finished"
                    finished-text="没有更多了"
                    @load="onLoad"
            >
                <van-panel
                        v-for="t in tasks"
                        :key="t.id"
                        :title="t.title"
                        :desc="t.description"
                        :status="t.status"
                        v-on:click="showTaskInfo(t)"/>
            </van-list>
        </van-pull-refresh>
    </div>
</template>
<script>
    export default {
        data() {
            return {
                loading: false,
                isLoading: false,
                finished: false,
                tasks: []
            }
        },
        beforeMount() {
            let cacheTasks = sessionStorage.getItem("tasks");
            if (cacheTasks !== null) {
                this.tasks = JSON.parse(cacheTasks);
            }
        },
        methods: {
            onRefresh: function () {
                let that = this;
                setTimeout(function () {
                    sessionStorage.setItem("tasks", "");
                    let newTasks = [];
                    for (let i = 0; i < 20; i++) {
                        newTasks.push({
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
                            }, {
                                url: "https://img.yzcdn.cn/vant/t-thirt.jpg",
                            }, {
                                url: "https://img.yzcdn.cn/vant/t-thirt.jpg",
                            }, {
                                url: "https://img.yzcdn.cn/vant/cat.jpeg",
                            }]
                        })
                    }
                    that.tasks = newTasks;
                    sessionStorage.setItem("tasks", JSON.stringify(newTasks));
                    that.$toast("success");
                    that.isLoading = false;
                    that.finished = false
                }, 1000);
            },
            onLoad: function () {
                let that = this;
                setTimeout(function () {
                    for (let i = 0; i < 20; i++) {
                        that.tasks.push({
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
                            }, {
                                url: "https://img.yzcdn.cn/vant/t-thirt.jpg",
                            }, {
                                url: "https://img.yzcdn.cn/vant/t-thirt.jpg",
                            }, {
                                url: "https://img.yzcdn.cn/vant/cat.jpeg",
                            }]
                        })
                    }
                    sessionStorage.setItem("tasks", JSON.stringify(that.tasks));
                    that.loading = false;
                    if (that.tasks.length >= 100) {
                        that.finished = true
                    }
                }, 1000);
            },
            showTaskInfo: function (task) {
                this.$router.push({
                    name: 'taskInfo',
                    params: {
                        task: task
                    },
                })
            },
            showAddTask: function () {
                this.$router.push('/addTask')
            }
        }
    }
</script>