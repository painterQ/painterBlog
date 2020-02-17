<template>
    <div class="painter-header">
        <div :style="logo">
            <svg width="297" height="100" viewBox="0 -10 594 200" xmlns="http://www.w3.org/2000/svg">
                <g>
                    <title>background</title>
                    <rect x="-1" y="-1" width="299" height="102" id="canvas_background" fill="#ffffff00"/>
                </g>
                <defs>
                    <filter height="200%" width="200%" y="-50%" x="-50%" id="svg_3_blur">
                        <feGaussianBlur stdDeviation="0" in="SourceGraphic"/>
                    </filter>
                </defs>
                <g>
                    <title>Layer 1</title>
                    <text filter="url(#svg_3_blur)" font-weight="bold" fill="#ffffff" stroke-width="0" x="58.20774" y="143.712482" id="svg_3" font-size="24" font-family="'Trebuchet MS', Gadget, sans-serif" text-anchor="start" xml:space="preserve" transform="matrix(1.923109531402588,0,0,2.672727346420288,-109.38847947120667,-311.1272864341736) " stroke="#000" font-style="normal">Painter's Blog</text>
                    <g stroke-width="0" fill="#ffffff" id="svg_5" transform="rotate(-14.649959564208984 65.00000762939453,19.000015258789062) ">
                        <path d="m78.38946,22.05328c4.22597,-1.48229 5.0617,-3.46048 4.21993,-6.97751c-0.95857,0.33495 -2.11953,-0.00109 -2.28768,-0.56334c-0.16916,-0.55789 0.58501,-1.56929 1.54358,-1.90642c-1.23145,-3.37892 -2.98446,-4.44686 -7.21044,-2.96566c-1.61708,0.56768 -3.26035,1.62258 -4.68713,2.86561c0.65247,-1.84987 1.07336,-3.86395 1.07336,-5.70077c0,-4.79813 -1.48317,-6.26954 -4.84019,-6.49031c0,1.08752 -0.6565,2.17504 -1.20123,2.17504s-1.20225,-1.08752 -1.20225,-2.17504c-3.35701,0.22077 -4.83918,1.69218 -4.83918,6.49031c0,1.83682 0.41988,3.8509 1.07337,5.69968c-1.42578,-1.24194 -3.07006,-2.29575 -4.68714,-2.86343c-4.22598,-1.48338 -5.97899,-0.41435 -7.21145,2.96566c0.95757,0.33713 1.71275,1.34635 1.5446,1.90642s-1.32912,0.89829 -2.28668,0.56333c-0.84278,3.51595 -0.00503,5.49414 4.21892,6.97643c1.61709,0.5666 3.52115,0.758 5.35069,0.65795c-1.53452,1.08099 -2.97037,2.44365 -3.97022,3.9292c-2.61191,3.88244 -2.21217,6.01398 0.38363,8.32387c0.59206,-0.87981 1.71476,-1.34309 2.15679,-0.99726c0.44102,0.34474 0.3806,1.64324 -0.21145,2.52304c2.83544,1.9521 4.83615,1.70197 7.44706,-2.1783c0.99985,-1.48554 1.75503,-3.38326 2.23331,-5.29295c0.47828,1.91077 1.23446,3.80741 2.23331,5.29295c2.61191,3.88244 4.61162,4.1304 7.44807,2.1783c-0.59106,-0.88089 -0.65349,-2.17721 -0.21146,-2.52413c0.44002,-0.34474 1.56372,0.11637 2.15679,0.99726c2.59479,-2.3088 2.99353,-4.44034 0.38363,-8.32278c-0.99985,-1.48555 -2.4367,-2.84821 -3.97022,-3.9292c1.82854,0.10005 3.73159,-0.09135 5.34968,-0.65795zm-13.3898,-0.51113c-0.16312,-0.31756 -0.35544,-0.58509 -0.59206,-0.77105c-0.23562,-0.18597 -0.53165,-0.30016 -0.86191,-0.37085c0.22957,-0.26535 0.40578,-0.54593 0.4964,-0.84609c0.09062,-0.29907 0.10069,-0.63837 0.05941,-0.99942c0.30408,0.15333 0.60615,0.24795 0.89816,0.24795s0.59407,-0.09462 0.90017,-0.24904c-0.04129,0.36105 -0.03122,0.70036 0.05941,0.99943c0.09061,0.30015 0.26682,0.58073 0.49539,0.84609c-0.32926,0.07068 -0.62529,0.18487 -0.8609,0.37084c-0.23763,0.18705 -0.43096,0.45458 -0.59407,0.77214z" fill="#ffffff" id="svg_1" stroke="null"/>
                    </g>
                </g>
            </svg>
        </div>
        <div class="header-profile">
            <el-form>
                <el-input placeholder="搜索..." prefix-icon="el-icon-search"></el-input>
            </el-form>

            <el-dropdown @command="logout">
                <el-avatar
                        shape="circle"
                        :size="40"
                        fit="cover"
                        :src="avatar"
                        @error="avatarError"
                ></el-avatar>
                <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item icon="el-icon-warning-outline" command="logout">退出登出</el-dropdown-item>
                </el-dropdown-menu>
            </el-dropdown>
        </div>
    </div>
</template>

<script>
    import api from '../api/rpc'
    import Vue from 'vue'
    import {Dropdown, DropdownItem, DropdownMenu} from "element-ui"
    import VueCookies from 'vue-cookies'
    import constVar from '../api/const'

    Vue.use(api);
    Vue.use(VueCookies);
    Vue.use(Dropdown);
    Vue.use(DropdownItem);
    Vue.use(DropdownMenu);
    export default {
        name: "index-body",
        data: function () {
            return {
                logo: {
                    height: "50px",
                    width: "150px",
                    margin: "0 1.5em 0 1.5em"
                },
            }
        },
        computed: {
            avatar: {
                get() {
                    return this.$store.state.author.avatar
                },
                set() {

                }
            },
        },
        methods: {
            avatarError() {
                return true
            },
            logout(command) {
                switch (command) {
                    case "logout":
                        this.$cookies.remove(constVar.cookieKey)
                        this.$store.commit("changeLogin", false)
                        break
                }
            }
        },
        async mounted() {
            let res = await this.$_getAuthorInfo()
            this.$store.commit("changeAvatar", res.data.avatar);
            this.$store.commit("changeName", res.data.name);
            this.$store.commit("changeMotto", res.data.say);
            this.$store.commit("changeIPC", res.data.ipc);
            this.$store.commit("changeTitle", res.data.title);
            this.$store.commit("changeSubTitle", res.data.subTitle);
            this.$store.commit("changeMail", res.data.email);
            this.$store.commit("changeGithub", res.data.github);
        }
    }
</script>

<style scoped>
    .painter-header {
        justify-content: space-between;
        display: flex;
        flex-wrap: wrap;
        /*learn: 渐变色，可以指定多个节点，可以追加百分比*/
        background: linear-gradient(to right, #545c64 20%, #fff);
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
        height: 50px;
        overflow-y: hidden;
    }

    .header-profile {
        width: 30%;
        display: flex;
    }

    .header-profile > * {
        margin: 5px 1em;
    }

</style>