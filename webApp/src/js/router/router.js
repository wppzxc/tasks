import Login from '../../pages/common/login'
import Register from '../../pages/common/register'
import Tasks from '../../pages/tasks/tasks'
import TasksInfo from '../../pages/tasks/taskInfo'
import AddTask from '../../pages/tasks/addTask'
import EditTask from '../../pages/tasks/editTask'
import Commits from '../../pages/commits/commits'
import CommitsInfo from '../../pages/commits/commitInfo'
import CommitSubmit from '../../pages/commits/commitSubmit'
import Settings from '../../pages/settings/settings'
import UserInfo from '../../pages/settings/userInfo'
import Service from '../../pages/settings/service'
import UpdateServiceInfo from '../../pages/settings/updateService'
import InviteInfo from '../../pages/settings/inviteInfo'

const routers = [
    {path: '/', redirect: '/taskInfo'},
    {path: '/login', name: "login", component: Login},
    {path: '/register/:parent', name: "register", component: Register},
    {path: '/tasks', name: "tasks", component: Tasks},
    {path: '/taskInfo', name: "taskInfo", component: TasksInfo},
    {path: '/addTask', name: "addTask", component: AddTask},
    {path: '/editTask', name: "editTask", component: EditTask},
    {path: '/commits', name: "commits", component: Commits},
    {path: '/commitInfo', name: "commitInfo", component: CommitsInfo},
    {path: '/commitSubmit', name: "commitSubmit", component: CommitSubmit},
    {path: '/settings', name: "settings", component: Settings},
    {path: '/userInfo', name: "userInfo", component: UserInfo},
    {path: '/service', name: "service", component: Service},
    {path: '/updateService', name: "updateService", component: UpdateServiceInfo},
    {path: '/inviteInfo', name: "inviteInfo", component: InviteInfo},
];

export default routers