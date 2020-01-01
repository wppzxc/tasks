<template>
    <div>
        <van-nav-bar
                title="任务详情"
                left-text="返回"
                left-arrow
                @click-left="goBack"/>
        <van-row>
            <van-col span="22" offset="1">
                <van-panel :title="task.title" :desc="task.description" :status="task.status">
                    <div v-text="task.longContent"></div>
                </van-panel>
            </van-col>
        </van-row>
        <van-row>
            <van-col span="8"
                     v-for="i in task.images">
                <van-image
                        width="100"
                        height="100"
                        fit="contain"
                        :src="i.url"
                />
            </van-col>
        </van-row>
        <van-row>
            <h3>{{comment}}</h3>
        </van-row>
        <van-row>
            <van-cell-group title=" ">
                <van-button type="info" v-on:click="onSubmit">完成</van-button>
            </van-cell-group>
        </van-row>
    </div>
</template>
<script>
    export default {
        data() {
            return {
                comment: "开始任务后才能查看评论内容哦~",
                task: {}
            }
        },
        mounted() {
            this.task = this.$route.params.task
        },
        methods: {
            onSubmit: function () {
                this.$router.push({
                    name: "commitSubmit",
                    params: {
                        task: this.task
                    },
                })
            },
            goBack: function () {
                this.$router.go(-1)
            }
        }
    }
</script>