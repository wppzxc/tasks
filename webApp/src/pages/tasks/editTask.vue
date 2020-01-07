<template>
    <div>
        <van-nav-bar title="修改任务"
                     left-text="返回"
                     left-arrow
                     @click-left="goBack"/>
        <van-cell title="基本信息"/>
        <van-cell-group>
            <van-field v-model="taskInfo.title" label="活动标题："/>
            <van-field
                    v-bind:value="taskInfo.expire | transTime" center clearable label="结束时间：">
                <van-button slot="button" size="mini" type="info" @click="showTimePicker=true">选择</van-button>
            </van-field>
            <van-field v-model="taskInfo.status" label="活动状态："/>
            <van-field v-model="taskInfo.bonus" label="任务赏金：￥" placeholder="单位是 分"/>
            <van-field v-model="taskInfo.parentBonus" label="一级赏金：￥" placeholder="单位是 分"/>
            <van-field v-model="taskInfo.grandParentBonus" label="二级赏金：￥" placeholder="单位是 分"/>
        </van-cell-group>
        <van-cell title="任务须知"/>
        <van-cell-group>
            <van-field
                    v-for="(s, key) in tmpSteps"
                    :key="key"
                    v-model="tmpSteps[key]"/>
        </van-cell-group>
        <van-button type="info" size="mini" @click="tmpSteps.push('')">添加</van-button>
        <van-cell title="详细教程"/>
        <van-uploader
                multiple
                v-model="tmpGuideImages"
                :max-count="5"/>
        <van-cell title="任务内容"/>
        <van-field v-model="taskInfo.commentBase" label="评论内容："/>
        <van-uploader
                multiple
                v-model="tmpCommentImages"
                :max-count="5"/>
        <van-row>
            <van-col offset="2" span="20">
                <van-button type="primary" round size="large" v-on:click="saveTaskInfo">保存</van-button>
            </van-col>
        </van-row>
        <van-popup
                title="请选择"
                v-model="showTimePicker"
                position="bottom"
                :style="{ height: '40%' }">
            <van-datetime-picker
                    v-model="taskInfo.expire"
                    @cancel="showTimePicker=false"
                    @confirm="confirmDatetime"
                    type="datetime"
                    :min-date="minDate"
                    :max-date="maxDate"
            />
        </van-popup>
    </div>
</template>
<script>
    import {BACK_HOST, TASKS, UPLOAD_IMAGES} from '../../js/const/const'
    import {mapState} from 'vuex'

    export default {
        data() {
            return {
                minHour: 10,
                maxHour: 20,
                minDate: new Date(),
                maxDate: new Date(2029, 12, 31),
                currentDate: new Date(),
                showTimePicker: false,
                taskInfo: {
                    ID: 0,
                    title: "",
                    status: "",
                    expire: "",
                    bonus: 0,
                    parentBonus: 0,
                    grandParentBonus: 0,
                    description: "",
                    guideImages: "",
                    commentBase: "",
                    commentImages: "",
                    steps: []
                },
                tmpSteps: [],
                tmpGuideImages: [],
                tmpCommentImages: [],
            }
        },
        computed: {
            ...mapState({
                user: state => state.user
            })
        },
        mounted() {
            this.initTasks();
        },
        methods: {
            confirmDatetime: function () {
                console.log(this.taskInfo.expire);
                this.showTimePicker = false
            },
            goBack: function () {
                this.$router.go(-1)
            },
            initTasks: function () {
                let that = this;
                that.$axios.get(BACK_HOST + that.user.name + TASKS).then(function (resp) {
                    console.log(resp.data);
                    let tmp = JSON.parse(resp.data.guideImages);
                    for (let i=0; i < tmp.length; i++) {
                        that.$axios.get(tmp[i], {responseType: "arraybuffer",}).then((resp)=>{
                            return  'data:image/png;base64,' + btoa(
                                new Uint8Array(resp.data).reduce((data, byte) => data + String.fromCharCode(byte), '')
                            );
                        }).then((data)=>{
                            that.tmpGuideImages[i]={
                                url: data,
                                isImage: true,
                            };
                        }).catch((err)=>{
                            console.log(err);
                            that.tmpGuideImages[i]={
                                url: "",
                                isImage: true,
                            };
                        });
                    }
                    let tmp2 = JSON.parse(resp.data.commentImages);
                    for (let i=0; i < tmp2.length; i++) {
                        that.$axios.get(tmp[i], {responseType: "arraybuffer",}).then((resp)=>{
                            return  'data:image/png;base64,' + btoa(
                                new Uint8Array(resp.data).reduce((data, byte) => data + String.fromCharCode(byte), '')
                            );
                        }).then((data)=>{
                            that.tmpCommentImages[i]={
                                url: data,
                                isImage: true,
                            };
                        }).catch((err)=>{
                            console.log(err);
                            that.tmpCommentImages[i]={
                                url: "",
                                isImage: true,
                            };
                        });
                    }
                    that.tmpSteps = JSON.parse(resp.data.steps);
                    that.taskInfo = resp.data
                }).catch(function (err) {
                    console.log(err);
                    that.$toast(err)
                })
            },
            saveTaskInfo: function () {
                let that = this;
                let formdata = new FormData();
                let allImg = that.tmpGuideImages.concat(that.tmpCommentImages);
                for (let i=0;i<allImg.length;i++) {
                    let key = "img" + i;
                    if (allImg[i].file === undefined) {
                        formdata.append(key, that.dataURLtoFile(allImg[i].url, key + ".png"));
                    } else {
                        formdata.append(key, allImg[i].file)
                    }
                }
                console.log(formdata);
                that.$axios.post(BACK_HOST + that.user.name + UPLOAD_IMAGES, formdata, {headers: {'Content-Type': 'multipart/form-data'}}).then((resp)=>{
                    that.taskInfo.guideImages = JSON.stringify(resp.data.slice(0, that.tmpGuideImages.length));
                    that.taskInfo.commentImages = JSON.stringify(resp.data.slice(that.tmpGuideImages.length));
                    that.taskInfo.steps = JSON.stringify(that.tmpSteps);
                    console.log(that.taskInfo);
                    that.$axios.put(BACK_HOST + that.user.name + TASKS, that.taskInfo).then((resp)=>{
                        console.log(resp);
                        that.goBack()
                    }).catch((err)=>{
                        console.log(err);
                        that.$toast(err)
                    })
                }).catch((err)=>{
                    that.$toast(err)
                });
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