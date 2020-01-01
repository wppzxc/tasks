<template>
    <div>
        <van-nav-bar title="任务详情">
            <van-icon name="edit" slot="right" v-on:click="showEditTask"/>
        </van-nav-bar>
        <van-panel :title="taskInfo.title" :desc="taskInfo.expire" :status="taskInfo.status">
            <div>
                <van-cell title="任务须知"/>
                <van-field
                        :border="false"
                        type="textarea"
                        disabled
                        autosize
                        v-for="(s, key) in steps"
                        :key="key"
                        :value="s"/>
            </div>
            <div>
                <van-cell title="详细步骤"/>
                <van-grid :column-num="3">
                    <van-grid-item v-for="(image, key) in taskInfo.guideImages" :key="key">
                        <van-image :src="image" @click="showImages(taskInfo.guideImages, key)"/>
                    </van-grid-item>
                </van-grid>
            </div>
            <div>
                <van-cell title="任务内容"/>
                <van-grid :column-num="3">
                    <van-grid-item v-for="(image, key) in taskInfo.commentImages" :key="key">
                        <van-image :src="image" @click="showImages(taskInfo.commentImages, key)"/>
                    </van-grid-item>
                </van-grid>
                <van-field label="评论内容："
                           placeholder="请开始任务，获取评论内容"
                           disabled
                           type="textarea"
                           autosize
                           v-model="currentCommit.comment"/>
            </div>
            <van-row>
                <van-col offset="2" span="20">
                    <van-button type="primary" round size="large">开始</van-button>
                </van-col>
            </van-row>
        </van-panel>
    </div>
</template>
<script>
    import {ImagePreview} from 'vant'
    export default {
        data() {
            return {
                comment: "开始任务后才能查看评论内容哦~",
                taskInfo: {
                    title: "评价有礼",
                    status: "进行中",
                    expire: "2020/1/30",
                    description: "1、打开饿了么，点击订单-找到等待评价的订单\n2、返回本页，复制评价，bao保存图片\n3、给商家骑手打好评，在评价里写网页li里的评价，上传网页里的图片，点击提交评价\n4、到商家页面，点击评价，点击最新，截图，回到网页上传，等待客服审核。任务审核完成，佣金直接到你的支付宝钱包",
                    guideImages: [
                        'https://img.yzcdn.cn/vant/apple-1.jpg',
                        'https://img.yzcdn.cn/vant/apple-2.jpg',
                        'https://img.yzcdn.cn/vant/apple-3.jpg',
                        'https://img.yzcdn.cn/vant/apple-4.jpg',
                        'https://img.yzcdn.cn/vant/apple-4.jpg',
                    ],
                    commentImages: [
                        "https://img.yzcdn.cn/vant/t-thirt.jpg",
                        "https://img.yzcdn.cn/vant/t-thirt.jpg",
                        "https://img.yzcdn.cn/vant/t-thirt.jpg",
                    ],
                    commentBase: "",
                },
                currentCommit: {
                    comment: ""
                },
                steps: ["1、打开饿了么，点击订单-找到等待评价的订单",
                    "2、返回本页，复制评价，bao保存图片",
                    "3、给商家骑手打好评，在评价里写网页li里的评价，上传网页里的图片，点击提交评价",
                    "4、到商家页面，点击评价，点击最新，截图，回到网页上传，等待客服审核。任务审核完成，佣金直接到你的支付宝钱包"],
            }
        },
        mounted() {
            this.task = this.$route.params.task
        },
        methods: {
            onStart: function () {
                let that = this;
                setTimeout(function () {
                    that.comment = "这件商品真是好啊！233"
                }, 1000)
            },
            onCommit: function () {
                this.$router.push({
                    name: "taskCommit",
                    params: {
                        task: this.task
                    },
                })
            },
            showImages: function(images, index) {
                ImagePreview({
                    images: images,
                    startPosition: index,
                });
            },
            goBack: function () {
                this.$router.go(-1)
            },
            showEditTask: function () {
                this.$router.push('/editTask')
            }
        }
    }
</script>
<style>
    .task-info {
        font-size: xx-large;
    }
</style>