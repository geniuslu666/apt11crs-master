<template>
  <div class="layout-header">
    <!--顶部菜单-->
    <audio ref="audio" controls hidden="true" :src="voicePath" />

    <!-- PC 端：顶部水平菜单 -->
    <div class="layout-header-left hidden md:flex"
      v-if="navMode === 'horizontal' || (navMode === 'horizontal-mix' && mixMenu)">
      <div class="logo" v-if="navMode === 'horizontal'">
        <img src="~@/assets/images/logo.png" alt="" />
        <h2 v-show="!collapsed" class="title">{{ projectName }}</h2>
      </div>
      <AsideMenu @update:collapsed="updateMenu" v-model:location="getMenuLocation" :inverted="getInverted"
        mode="horizontal" />
    </div>

    <!-- PC 端：左侧菜单模式 -->
    <div class="layout-header-left hidden md:flex" v-else>
      <!-- 菜单收起 -->
      <div class="ml-1 layout-header-trigger layout-header-trigger-min"
        @click="() => $emit('update:collapsed', !collapsed)">
        <n-icon size="18" v-if="collapsed">
          <MenuUnfoldOutlined />
        </n-icon>
        <n-icon size="18" v-else>
          <MenuFoldOutlined />
        </n-icon>
      </div>
      <!-- 刷新 -->
      <div class="mr-1 layout-header-trigger layout-header-trigger-min" v-if="headerSetting.isReload"
        @click="reloadPage">
        <n-icon size="18">
          <ReloadOutlined />
        </n-icon>
      </div>
      <!-- 面包屑 -->
      <n-breadcrumb v-if="crumbsSetting.show">
        <template v-for="routeItem in breadcrumbList" :key="routeItem.name">
          <n-breadcrumb-item>
            <n-dropdown v-if="routeItem.children.length" :options="routeItem.children" @select="dropdownSelect">
              <span class="link-text">
                <component v-if="crumbsSetting.showIcon && routeItem.meta.icon" :is="routeItem.meta.icon" />
                {{ routeItem.meta.title }}
              </span>
            </n-dropdown>
            <span class="link-text" v-else>
              <component v-if="crumbsSetting.showIcon && routeItem.meta.icon" :is="routeItem.meta.icon" />
              {{ routeItem.meta.title }}
            </span>
          </n-breadcrumb-item>
        </template>
      </n-breadcrumb>
    </div>

    <!-- PC 端：右侧功能区 -->
    <div class="layout-header-right hidden md:flex">
      <!--      <div-->
      <!--        class="layout-header-trigger layout-header-trigger-min"-->
      <!--        v-for="item in iconList"-->
      <!--        :key="item.icon.name"-->
      <!--      >-->
      <!--        <n-tooltip placement="bottom">-->
      <!--          <template #trigger>-->
      <!--            <n-icon size="18">-->
      <!--              <component :is="item.icon" v-on="item.eventObject || {}" />-->
      <!--            </n-icon>-->
      <!--          </template>-->
      <!--          <span>{{ item.tips }}</span>-->
      <!--        </n-tooltip>-->
      <!--      </div>-->

      <div class="layout-header-trigger layout-header-trigger-min" v-for="item in iconList" :key="item.icon.name">
        <n-popover placement="bottom" v-if="item.icon === 'BellOutlined'" trigger="click"
          :width="getIsMobile ? 276 : 420">
          <template #trigger>
            <n-tooltip placement="bottom">
              <template #trigger>
                <n-badge :value="notificationStore.getUnreadCount()" :max="99" processing>
                  <n-icon size="18">
                    <BellOutlined />
                  </n-icon>
                </n-badge>
              </template>
              <span>{{ item.tips }}</span>
            </n-tooltip>
          </template>

          <SystemMessage />
        </n-popover>

        <div v-else>
          <n-tooltip placement="bottom">
            <template #trigger>
              <n-icon size="18">
                <component :is="item.icon" v-on="item.eventObject || {}" />
              </n-icon>
            </template>
            <span>{{ item.tips }}</span>
          </n-tooltip>
        </div>
      </div>
      <!--切换全屏-->
      <div class="layout-header-trigger layout-header-trigger-min">
        <a-dropdown>
          <a class="ant-dropdown-link" @click.prevent style="color: rgb(51, 54, 57)">
            <n-icon size="18">
              <GlobeOutline />
            </n-icon>
            <!--            <DownOutlined />-->
          </a>
          <template #overlay>
            <a-menu>
              <a-menu-item>
                <a @click="changelang('zh')" :class="userStore.language == 'zh' ? 'cblue' : ''">简体中文</a>
              </a-menu-item>
              <a-menu-item>
                <a @click="changelang('ja')" :class="userStore.language == 'ja' ? 'cblue' : ''">日本语
                </a>
              </a-menu-item>
              <a-menu-item>
                <a @click="changelang('en')" :class="userStore.language == 'en' ? 'cblue' : ''">English
                </a>
              </a-menu-item>
              <a-menu-item>
                <a @click="changelang('ko')" :class="userStore.language == 'ko' ? 'cblue' : ''">
                  한국어
                </a>
              </a-menu-item>
              <a-menu-item>
                <a @click="changelang('zh_CN')" :class="userStore.language == 'zh_CN' ? 'cblue' : ''">
                  繁体中文
                </a>
              </a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>
      </div>
      <!--切换全屏-->
      <div class="layout-header-trigger layout-header-trigger-min">
        <n-tooltip placement="bottom">
          <template #trigger>
            <n-icon size="18">
              <component :is="fullscreenIcon" @click="toggleFullScreen" />
            </n-icon>
          </template>
          <span>全屏</span>
        </n-tooltip>
      </div>
      <!-- 个人中心 -->
      <div class="layout-header-trigger layout-header-trigger-min">
        <n-dropdown trigger="click" @select="avatarSelect" :options="avatarOptions" show-arrow>
          <div class="avatar">
            <n-avatar v-if="userStore.avatar" round :size="30" :src="userStore.avatar" />
            <n-avatar v-else round :size="30">{{ userStore.realName }}</n-avatar>
          </div>
        </n-dropdown>
      </div>
      <!--设置-->
      <!-- <div class="layout-header-trigger layout-header-trigger-min" @click="openSetting">
        <n-tooltip placement="bottom-end">
          <template #trigger>
            <n-icon size="18" style="font-weight: bold">
              <SettingOutlined />
            </n-icon>
          </template>
          <span>项目配置</span>
        </n-tooltip>
      </div> -->
    </div>

    <!-- 移动端：Header -->
    <div class="mobile-header w-full flex md:hidden items-center justify-between px-4 h-12">
      <!-- 左侧：一级 / 二级菜单按钮组 -->
      <div class="flex items-center gap-1">
        <!-- 一级导航（系统级） -->
        <button
          class="mobile-menu-btn w-10 h-10 flex items-center justify-center rounded-md hover:bg-gray-100 active:bg-gray-200 transition-colors"
          @click="toggleMobileDrawer" aria-label="打开一级菜单">
          <n-icon size="20">
            <MenuUnfoldOutlined />
          </n-icon>
        </button>

        <!-- 二级菜单（当前系统业务菜单） -->
        <button
          class="mobile-menu-btn w-10 h-10 flex items-center justify-center rounded-md hover:bg-gray-100 active:bg-gray-200 transition-colors"
          @click="$emit('toggle-mobile-aside')" aria-label="打开二级菜单">
          <n-icon size="24">
            <ListOutline />
          </n-icon>
        </button>
      </div>

      <!-- 中间：系统名称 -->
      <div class="mobile-title text-base font-semibold truncate px-2 flex-1 text-center cursor-pointer" @click="goHome">
        {{ projectName }}
      </div>

      <!-- 右侧：常用功能（移动端不显示消息通知，改为多语言切换） -->
      <div class="mobile-actions flex items-center gap-1">
        <!-- 消息通知（移动端暂不需要，保留代码方便以后启用）
        <div
          class="mobile-action-btn w-10 h-10 flex items-center justify-center rounded-md hover:bg-gray-100 active:bg-gray-200 transition-colors">
          <n-popover placement="bottom-end" trigger="click" :width="276">
            <template #trigger>
              <n-badge :value="notificationStore.getUnreadCount()" :max="99" processing>
                <n-icon size="22">
                  <BellOutlined />
                </n-icon>
              </n-badge>
            </template>
            <SystemMessage />
          </n-popover>
        </div>
        -->

        <!-- 多语言切换 -->
        <div
          class="mobile-action-btn w-10 h-10 flex items-center justify-center rounded-md hover:bg-gray-100 active:bg-gray-200 transition-colors">
          <a-dropdown>
            <a class="ant-dropdown-link" @click.prevent style="color: rgb(51, 54, 57)">
              <n-icon size="20">
                <GlobeOutline />
              </n-icon>
            </a>
            <template #overlay>
              <a-menu>
                <a-menu-item>
                  <a @click="changelang('zh')" :class="userStore.language == 'zh' ? 'cblue' : ''">简体中文</a>
                </a-menu-item>
                <a-menu-item>
                  <a @click="changelang('ja')" :class="userStore.language == 'ja' ? 'cblue' : ''">日本语</a>
                </a-menu-item>
                <a-menu-item>
                  <a @click="changelang('en')" :class="userStore.language == 'en' ? 'cblue' : ''">English</a>
                </a-menu-item>
                <a-menu-item>
                  <a @click="changelang('ko')" :class="userStore.language == 'ko' ? 'cblue' : ''">한국어</a>
                </a-menu-item>
                <a-menu-item>
                  <a @click="changelang('zh_CN')" :class="userStore.language == 'zh_CN' ? 'cblue' : ''">
                    繁体中文
                  </a>
                </a-menu-item>
              </a-menu>
            </template>
          </a-dropdown>
        </div>

        <!-- 个人中心 -->
        <div
          class="mobile-action-btn w-10 h-10 flex items-center justify-center rounded-md hover:bg-gray-100 active:bg-gray-200 transition-colors">
          <n-dropdown trigger="click" @select="avatarSelect" :options="avatarOptions" show-arrow>
            <div class="flex items-center justify-center">
              <n-avatar v-if="userStore.avatar" round :size="28" :src="userStore.avatar" />
              <n-avatar v-else round :size="28">{{ userStore.realName }}</n-avatar>
            </div>
          </n-dropdown>
        </div>
      </div>
    </div>

    <!-- 移动端：导航抽屉（一级菜单），宽度与二级菜单保持一致 -->
    <n-drawer v-model:show="mobileDrawerVisible" :width="menuWidth" :placement="'left'" class="mobile-nav-drawer"
      :mask-closable="true">
      <template #header>
        <div class="flex items-center justify-between px-4 h-14 border-b">
          <div class="text-base font-semibold">{{ projectName }}</div>
          <button
            class="w-8 h-8 flex items-center justify-center rounded-md hover:bg-gray-100 active:bg-gray-200 transition-colors"
            @click="mobileDrawerVisible = false" aria-label="关闭菜单">
            <n-icon size="18">
              <MenuFoldOutlined />
            </n-icon>
          </button>
        </div>
      </template>
      <div class="mobile-nav-content">
        <AsideMenu @clickMenuItem="handleMobileMenuClick" v-model:location="getMenuLocation" :inverted="getInverted" />
      </div>
    </n-drawer>
  </div>
  <!--项目配置-->
  <ProjectSetting ref="drawerSetting" />

  <pmsAppReservationView ref="pmsAppReservationViewRef" />
  <hisorderview ref="hisorderviewRef" />

  <template>
    <n-notification-provider :max="3" />
  </template>
</template>

<script lang="ts">
import {
  defineComponent,
  reactive,
  toRefs,
  ref,
  computed,
  unref,
  watch,
  h,
  onMounted,
  nextTick,
} from 'vue';
import { useRouter, useRoute } from 'vue-router';
import pmsAppReservationView from '@/views/pmsAppReservation/view.vue';
import hisorderview from '@/views/pmsAppReservation/hisorderview.vue';

import components from './components';
import {
  NDialogProvider,
  useDialog,
  useMessage,
  NAvatar,
  NTag,
  NIcon,
  useNotification,
  NotificationReactive,
  NButton,
  NText,
} from 'naive-ui';
import { TABS_ROUTES } from '@/store/mutation-types';
import { useUserStore } from '@/store/modules/user';
import { useLockscreenStore } from '@/store/modules/lockscreen';
import ProjectSetting from './ProjectSetting.vue';
import { AsideMenu } from '@/layout/components/Menu';
import { useProjectSetting } from '@/hooks/setting/useProjectSetting';
import { NotificationsOutline as NotificationsIcon, GlobeOutline, ListOutline } from '@vicons/ionicons5';
import { DownOutlined } from '@vicons/antd';
import SystemMessage from './SystemMessage.vue';
import { notificationStoreWidthOut } from '@/store/modules/notification';
import { getIcon } from '@/enums/systemMessageEnum';

import { lastOrder } from '@/api/pmsRoomReservation';

export default defineComponent({
  name: 'PageHeader',
  components: {
    ...components,
    NDialogProvider,
    ProjectSetting,
    AsideMenu,
    SystemMessage,
    GlobeOutline,
    pmsAppReservationView,
    hisorderview,
    DownOutlined,
    ListOutline,
  },
  emits: ['update:collapsed', 'toggle-mobile-aside'],
  props: {
    collapsed: {
      type: Boolean,
    },
    inverted: {
      type: Boolean,
    },
  },
  setup(props, { emit }) {
    const userStore = useUserStore();
    const notificationStore = notificationStoreWidthOut();
    const useLockscreen = useLockscreenStore();
    const message = useMessage();
    const dialog = useDialog();
    const {
      getNavMode,
      getNavTheme,
      getHeaderSetting,
      getMenuSetting,
      getCrumbsSetting,
      getIsMobile,
    } = useProjectSetting();

    // const { username, avatar } = userStore?.info || {};
    const drawerSetting = ref();
    const pmsAppReservationViewRef = ref();
    const hisorderviewRef = ref();

    const projectName = import.meta.env.VITE_GLOB_APP_TITLE;
    const audio = ref();
    const voicePath = ref('');

    // 移动端抽屉显示状态
    const mobileDrawerVisible = ref(false);

    const state = reactive({
      // username: username || '',
      // avatar: avatar || '',
      fullscreenIcon: 'FullscreenOutlined',
      navMode: getNavMode,
      navTheme: getNavTheme,
      headerSetting: getHeaderSetting,
      crumbsSetting: getCrumbsSetting,
    });

    // 抽屉宽度：与 PC 侧边栏 menuWidth 保持一致，避免一二级宽度不一致
    const menuWidth = computed(() => {
      const { menuWidth } = unref(getMenuSetting) as any;
      return menuWidth;
    });

    const getInverted = computed(() => {
      const navTheme = unref(getNavTheme);
      return ['light', 'header-dark'].includes(navTheme) ? props.inverted : !props.inverted;
    });

    const mixMenu = computed(() => {
      return unref(getMenuSetting).mixMenu;
    });

    const getChangeStyle = computed(() => {
      const { collapsed } = props;
      const { minMenuWidth, menuWidth }: any = unref(getMenuSetting);
      return {
        left: collapsed ? `${minMenuWidth}px` : `${menuWidth}px`,
        width: `calc(100% - ${collapsed ? `${minMenuWidth}px` : `${menuWidth}px`})`,
      };
    });

    const playVoice = () => {
      // 外部链接
      //voicePath.value = `https://mp3在线地址`
      // 本地链接
      voicePath.value = new URL('@/assets/images/alert.mp3', import.meta.url).href;

      nextTick(() => {
        // 从头开始
        audio.value.currentTime = 0;
        // 播放
        audio.value?.play();
      });
    };
    const getMenuLocation = computed(() => {
      return 'header';
    });

    const router = useRouter();
    const route = useRoute();
    const generator: any = (routerMap) => {
      return routerMap.map((item) => {
        const currentMenu = {
          ...item,
          label: item.meta.title,
          key: item.name,
          disabled: item.path === '/',
        };
        // 是否有子菜单，并递归处理
        if (item.children && item.children.length > 0) {
          // Recursion
          currentMenu.children = generator(item.children, currentMenu);
        }
        return currentMenu;
      });
    };

    const breadcrumbList = computed(() => {
      return generator(route.matched);
    });

    const dropdownSelect = (key) => {
      router.push({ name: key });
    };

    // 刷新页面
    const reloadPage = () => {
      const full = unref(route);
      router.push({
        path: '/redirect' + full.path,
        query: full.query,
      });
    };

    // 注销登录
    const doLogout = () => {
      dialog.info({
        title: '提示',
        content: '您确定要注销登录吗',
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: () => {
          userStore.logout().then(() => {
            message.success('成功注销登录');
            // 移除标签页
            localStorage.removeItem(TABS_ROUTES);
            router
              .replace({
                name: 'Login',
                query: {
                  redirect: route.fullPath,
                },
              })
              .finally(() => location.reload());
          });
        },
        onNegativeClick: () => { },
      });
    };

    // 切换全屏图标
    const toggleFullscreenIcon = () =>
    (state.fullscreenIcon =
      document.fullscreenElement !== null ? 'FullscreenExitOutlined' : 'FullscreenOutlined');

    // 监听全屏切换事件
    document.addEventListener('fullscreenchange', toggleFullscreenIcon);

    // 全屏切换
    const toggleFullScreen = () => {
      if (!document.fullscreenElement) {
        document.documentElement.requestFullscreen();
      } else {
        if (document.exitFullscreen) {
          document.exitFullscreen();
        }
      }
    };

    // 图标列表
    const iconList = [
      // {
      //   icon: 'SearchOutlined',
      //   tips: '搜索',
      // },
      // {
      //   icon: 'GithubOutlined',
      //   tips: 'github',
      //   eventObject: {
      //     click: () => window.open('https://github.com/bufanyun/hotgo'),
      //   },
      // },
      // {
      //   icon: 'BellOutlined',
      //   tips: '我的消息',
      // },
      // {
      //   icon: 'LockOutlined',
      //   tips: '锁屏',
      //   eventObject: {
      //     click: () => useLockscreen.setLock(true),
      //   },
      // },
    ];
    const changelang = (type) => {
      userStore.changelang(type);
      window.location.reload();
    };
    function renderCustomHeader() {
      return h(
        'div',
        {
          style: 'display: flex; align-items: center; padding: 8px 12px;',
        },
        [
          h('div', null, [
            h('div', null, [
              h(NText, { depth: 2 }, { default: () => userStore?.info?.username }),
            ]),
            h('div', { style: 'font-size: 12px;' }, [
              h(NText, { depth: 3 }, { default: () => userStore?.info?.roleName }),
            ]),
          ]),
        ]
      );
    }

    const avatarOptions = [
      {
        key: 'header',
        type: 'render',
        render: renderCustomHeader,
      },
      {
        type: 'divider',
        key: 'd1',
      },
      {
        label: '个人设置',
        key: 1,
      },
      {
        label: '注销登录',
        key: 2,
      },
    ];

    //头像下拉菜单
    const avatarSelect = (key) => {
      switch (key) {
        case 1:
          router.push({ name: 'home_account' });
          break;
        case 2:
          doLogout();
          break;
      }
    };

    function openSetting() {
      const { openDrawer } = drawerSetting.value;
      openDrawer();
    }

    const notification = useNotification();
    const getMessages = computed(() => {
      return notificationStore.newMessage;
    });
    const nRef = ref<NotificationReactive | null>(null);
    // 监听新消息，推送通知
    watch(
      getMessages,
      async (newVal, _oldVal) => {
        //订单方面
        if (newVal && newVal.type == 4) {
          let content = JSON.parse(newVal.content);
          // const memberId = content.to.split('|')[1];
          // const res = await lastOrder({ MemberId: memberId });
          // if (!res.orderSn) {
          //   message.error('暂无订单信息');
          //   return;
          // }
          // content.orderSn = res.orderSn;
          // content.memberId = memberId;
          if (content.type == 'orderList') {
            pmsAppReservationViewRef.value.openModal({
              viewtype: '订单详情',
              content: content.content,
            });
          }
          // if (content.type == 'orderList_his') {
          //   hisorderviewRef.value.openModal(memberId);
          // }
        } else {
          if (newVal === null || newVal === undefined) {
            return;
          }
          if (newVal.title == '客服新消息通知') {
            playVoice();
          }
          nRef.value = notification.create({
            title: newVal.title,
            duration: 5000,
            description:
              newVal.tagTitle === '' || newVal.tagTitle === undefined
                ? undefined
                : () =>
                  h(
                    NTag,
                    {
                      style: {
                        marginRight: '6px',
                      },
                      type: newVal.tagProps?.type,
                      bordered: false,
                    },
                    {
                      default: () => newVal.tagTitle,
                    }
                  ),

            content: () =>
              newVal.content === '' || newVal.content === undefined
                ? undefined
                : h('div', { innerHTML: '<div>' + newVal.content + '</div>' }),
            meta: newVal.createdAt,
            avatar: () =>
              newVal.senderAvatar !== '' || newVal.senderAvatar === undefined
                ? h(NAvatar, {
                  size: 'small',
                  round: true,
                  src: newVal.senderAvatar,
                })
                : h(NIcon, null, { default: () => h(getIcon(newVal)) }),
            action: () =>
              h(
                NButton,
                {
                  text: true,
                  type: 'info',
                  onClick: () => {
                    (nRef.value as NotificationReactive).destroy();

                    if (newVal.title !== '客服新消息通知') {
                      router.push({
                        name: 'home_message',
                        query: {
                          type: newVal.type,
                        },
                      });
                    } else {
                      router.push({
                        name: 'kefuCenter',
                        query: {
                          key: 'chat_main',
                        },
                      });
                    }
                  },
                },
                {
                  default: () => '查看详情',
                }
              ),
            onClose: () => {
              nRef.value = null;
            },
          });
        }
      },
      { immediate: true, deep: true }
    );

    const updateMenu = () => {
      emit('update:collapsed', !props.collapsed);
    };

    // 切换移动端抽屉
    const toggleMobileDrawer = () => {
      mobileDrawerVisible.value = !mobileDrawerVisible.value;
    };

    // 移动端菜单点击后关闭抽屉
    const handleMobileMenuClick = (key: string) => {
      mobileDrawerVisible.value = false;
      // 触发原有的菜单点击逻辑
      if (/http(s)?:/.test(key)) {
        window.open(key);
      } else {
        router.push({ name: key });
      }
    };

    // 返回首页
    const goHome = () => {
      router.push({ path: '/' });
    };

    onMounted(() => {
      if (notificationStore.getUnreadCount() === 0) {
        notificationStore.pullMessages();
      }
    });
    return {
      ...toRefs(state),
      iconList,
      toggleFullScreen,
      doLogout,
      route,
      dropdownSelect,
      avatarOptions,
      getChangeStyle,
      avatarSelect,
      breadcrumbList,
      changelang,
      reloadPage,
      drawerSetting,
      pmsAppReservationViewRef,
      hisorderviewRef,
      voicePath,
      audio,
      openSetting,
      getInverted,
      getMenuLocation,
      mixMenu,
      NotificationsIcon,
      SystemMessage,
      notificationStore,
      getIsMobile,
      userStore,
      updateMenu,
      projectName,
      menuWidth,
      mobileDrawerVisible,
      toggleMobileDrawer,
      handleMobileMenuClick,
      goHome,
    };
  },
});
</script>

<style lang="less" scoped>
// Stripe-style PC header
.layout-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0;
  height: @header-height;
  background: #ffffff;
  border-bottom: 1px solid #e3e8ef;
  box-shadow: none;
  transition: border-color 0.15s;
  width: 100%;
  z-index: 11;

  &-left {
    align-items: center;

    .logo {
      display: flex;
      align-items: center;
      justify-content: center;
      height: @header-height;
      line-height: @header-height;
      overflow: hidden;
      white-space: nowrap;
      padding-left: 16px;
      min-width: 200px;

      img {
        width: auto;
        height: 28px;
        margin-right: 10px;
      }

      .title {
        margin-bottom: 0;
        min-width: 120px;
        font-size: 14px;
        font-weight: 600;
        color: #1a1f36;
        letter-spacing: -0.02em;
      }
    }

    ::v-deep(.ant-breadcrumb span:last-child .link-text) {
      color: #697386;
    }

    .n-breadcrumb {
      display: inline-block;
    }

    ::v-deep(.n-breadcrumb-item__link) {
      color: #697386;
      font-size: 13px;

      &:hover {
        color: #1a1f36;
      }
    }

    &-menu {
      color: var(--text-color);
    }
  }

  &-right {
    align-items: center;
    margin-right: 16px;
    gap: 4px;

    .avatar {
      display: flex;
      align-items: center;
      height: @header-height;
    }

    >* {
      cursor: pointer;
    }
  }

  &-trigger {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 36px;
    height: 36px;
    border-radius: 6px;
    cursor: pointer;
    transition: background 0.15s;
    color: #697386;

    .n-icon {
      display: flex;
      align-items: center;
    }

    &:hover {
      background: #f0f3f7;
      color: #1a1f36;
    }

    .anticon {
      font-size: 16px;
    }
  }

  &-trigger-min {
    width: auto;
    padding: 0 8px;
  }
}

.layout-header-light {
  background: #ffffff;
  color: #1a1f36;

  .n-icon {
    color: #697386;
  }

  .layout-header-left {
    ::v-deep(.n-breadcrumb .n-breadcrumb-item:last-child .n-breadcrumb-item__link) {
      color: #697386;
    }
  }

  .layout-header-trigger {
    &:hover {
      background: #f0f3f7;
    }
  }
}

.layout-header-fix {
  position: fixed;
  top: 0;
  right: 0;
  left: 200px;
  z-index: 11;
}

::v-deep(.menu-server-link) {
  color: #697386;
  font-size: 13px;
  font-weight: 500;

  &:hover {
    color: #635bff;
  }
}

:deep(sup) {
  top: 1.3em;
}

.cblue {
  color: #635bff !important;
}

/* 移动端 Header 样式 */
.mobile-header {
  background: #ffffff;
  border-bottom: 1px solid #e3e8ef;
  box-shadow: none;
  z-index: 11;
}

.mobile-menu-btn,
.mobile-action-btn {
  color: #697386;
  -webkit-tap-highlight-color: transparent;

  &:hover {
    color: #1a1f36;
  }
}

.mobile-title {
  color: #1a1f36;
  font-weight: 600;
  font-size: 15px;
  letter-spacing: -0.02em;
}

/* 移动端导航抽屉样式 */
.mobile-nav-drawer {
  :deep(.n-drawer-body-content-wrapper) {
    padding: 0;
    background: #0a2540;
  }

  :deep(.n-drawer-header) {
    background: #0a2540;
    border-bottom: 1px solid rgba(255,255,255,0.08);
    color: #ffffff;
    padding: 0 16px;
    height: 56px;
  }
}

.mobile-nav-content {
  padding: 8px 0;
  overflow-y: auto;
  height: calc(100vh - 56px);
  background: #0a2540;

  :deep(.n-menu) {
    background-color: #0a2540;

    .n-menu-item-content {
      color: #c8d2e0;
      font-size: 13px;
      font-weight: 500;
      min-height: 40px;
      border-radius: 6px;
      margin: 1px 8px;
    }

    .n-menu-item-content:hover {
      color: #ffffff;
      background-color: rgba(99, 91, 255, 0.15);
    }

    .n-menu-item-content--selected {
      color: #ffffff;
      background-color: rgba(99, 91, 255, 0.25);
    }

    .n-submenu-children .n-menu-item-content {
      padding-left: 28px;
      min-height: 36px;
    }
  }
}
</style>
