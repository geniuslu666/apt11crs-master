<template>
  <n-layout class="layout" :position="fixedMenu" has-sider>
    <n-layout-sider v-if="!isMobile" :position="fixedMenu"
      :collapsed="collapsed" collapse-mode="width" :collapsed-width="64" :width="leftMenuWidth"
      :native-scrollbar="false" :inverted="false" class="layout-sider">
      <div class="sidebar-inner" :class="{ 'sidebar-inner-collapsed': collapsed }">
        <Logo :collapsed="collapsed" />
        <div class="sidebar-menu-wrap">
          <AsideMenu v-model:collapsed="collapsed" v-model:location="getMenuLocation" />
        </div>
        <div class="sidebar-footer">
          <button class="sidebar-collapse-btn" type="button" @click="collapsed = !collapsed" :title="collapsed ? '展开菜单' : '折叠菜单'">
            <svg v-if="collapsed" xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m9 18 6-6-6-6"/></svg>
            <svg v-else xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m15 18-6-6 6-6"/></svg>
          </button>
        </div>
      </div>
    </n-layout-sider>

    <!-- 移动端：业务侧边二级菜单抽屉（与顶部一级菜单分离） -->
    <n-drawer v-if="isMobile" v-model:show="mobileAsideVisible" :width="menuWidth" placement="left"
      class="layout-side-drawer md:hidden" :mask-closable="true">
      <Logo :collapsed="collapsed" />
      <AsideMenu :inverted="false" @clickMenuItem="handleMobileAsideClick" />
    </n-drawer>

    <n-layout :inverted="inverted">
      <n-layout-header :inverted="getHeaderInverted" :position="fixedHeader">
        <PageHeader v-model:collapsed="collapsed" :inverted="inverted" @toggle-mobile-aside="toggleMobileAside" />
      </n-layout-header>

      <n-layout-content class="layout-content" :class="{ 'layout-default-background': getDarkTheme === false }">
        <div class="layout-content-main" :class="{
          'layout-content-main-fix': fixedMulti,
          'fluid-header': fixedHeader === 'static',
        }">
          <TabsView v-if="isMultiTabs" v-model:collapsed="collapsed" />
          <div class="main-view" :class="{
            'main-view-fix': fixedMulti,
            noMultiTabs: !isMultiTabs,
            'mt-3': !isMultiTabs,
          }">
            <MainView />
          </div>
        </div>
      </n-layout-content>
      <n-back-top :right="100" />
    </n-layout>
  </n-layout>
</template>

<script lang="ts" setup>
import { ref, unref, computed, onMounted } from 'vue';
import { Logo } from './components/Logo';
import { TabsView } from './components/TagsView';
import { MainView } from './components/Main';
import { AsideMenu } from './components/Menu';
import { PageHeader } from './components/Header';
import { useProjectSetting } from '@/hooks/setting/useProjectSetting';
import { useDesignSetting } from '@/hooks/setting/useDesignSetting';
import { useLoadingBar } from 'naive-ui';
import { useRoute } from 'vue-router';
import { useProjectSettingStore } from '@/store/modules/projectSetting';

const { getDarkTheme } = useDesignSetting();
const {
  // getShowFooter,
  getNavTheme,
  getHeaderSetting,
  getMenuSetting,
  getMultiTabsSetting,
} = useProjectSetting();

const route = useRoute();
const settingStore = useProjectSettingStore();

const navMode = computed(() => 'vertical');

const collapsed = ref<boolean>(false);

const { mobileWidth, menuWidth } = unref(getMenuSetting);

const isMobile = computed<boolean>({
  get: () => settingStore.getIsMobile,
  set: (val) => settingStore.setIsMobile(val),
});

const fixedHeader = computed(() => {
  const { fixed } = unref(getHeaderSetting);
  return fixed ? 'absolute' : 'static';
});


const isMixMenuNoneSub = computed(() => {
  const mixMenu = settingStore.menuSetting.mixMenu;
  // const currentRoute = useRoute();
  if (unref(navMode) != 'horizontal-mix') return true;
  if (unref(navMode) === 'horizontal-mix' && mixMenu && route.meta.isRoot) {
    return false;
  }
  return true;
});

const fixedMenu = computed(() => {
  const { fixed } = unref(getHeaderSetting);
  return fixed ? 'absolute' : 'static';
});

const isMultiTabs = computed(() => {
  return unref(getMultiTabsSetting).show;
});

const fixedMulti = computed(() => {
  return unref(getMultiTabsSetting).fixed;
});

const inverted = computed(() => {
  return ['dark', 'header-dark'].includes(unref(getNavTheme));
});

const getHeaderInverted = computed(() => {
  const navTheme = unref(getNavTheme);
  return ['light', 'header-dark'].includes(navTheme) ? unref(inverted) : !unref(inverted);
});

const leftMenuWidth = computed(() => {
  return collapsed.value ? 64 : 224;
});

// const getChangeStyle = computed(() => {
//   const { minMenuWidth, menuWidth } = unref(getMenuSetting);
//   return {
//     'padding-left': collapsed.value ? `${minMenuWidth}px` : `${menuWidth}px`,
//   };
// });

const getMenuLocation = computed(() => {
  return 'left';
});

// 移动端业务侧边二级菜单抽屉可见性（与顶部一级导航分离）
const mobileAsideVisible = ref(false);

const toggleMobileAside = () => {
  if (!isMobile.value) return;
  mobileAsideVisible.value = !mobileAsideVisible.value;
};

const handleMobileAsideClick = () => {
  mobileAsideVisible.value = false;
};

//判断是否触发移动端模式
const checkMobileMode = () => {
  if (document.body.clientWidth <= mobileWidth) {
    isMobile.value = true;
  } else {
    isMobile.value = false;
  }
  collapsed.value = false;
};

const watchWidth = () => {
  const Width = document.body.clientWidth;
  if (Width <= 950) {
    collapsed.value = true;
  } else collapsed.value = false;

  checkMobileMode();
};

onMounted(() => {
  checkMobileMode();
  window.addEventListener('resize', watchWidth);
  //挂载在 window 方便与在js中使用
  window['$loading'] = useLoadingBar();
  window['$loading'].finish();
});
</script>

<style lang="less">
// Mobile sidebar drawer
.layout-side-drawer {
  background-color: #ffffff !important;

  :deep(.n-drawer-body-content-wrapper) {
    padding: 0;
    background-color: #ffffff;
  }
}
</style>
<style lang="less" scoped>
.layout {
  display: flex;
  flex-direction: row;
  flex: auto;

  &-default-background {
    background: #f7fbff;
  }

  .layout-sider {
    min-height: 100vh;
    box-shadow: none;
    border-right: 1px solid hsl(214 32% 91%);
    position: relative;
    z-index: 13;
    transition: width 0.2s ease;
    background: hsl(0 0% 100%) !important;
    padding: 0;

    :deep(.n-layout-sider-scroll-container) {
      background: transparent;
      scrollbar-width: none;

      &::-webkit-scrollbar {
        width: 0;
        height: 0;
      }
    }
  }

  .layout-content {
    flex: auto;
    min-height: 100vh;
    background: #f7fbff;
  }

  .n-layout-header.n-layout-header--absolute-positioned {
    z-index: 11;
  }

  .n-layout-footer {
    background: none;
  }
}

// Sidebar flex wrapper
.sidebar-inner {
  display: flex;
  flex-direction: column;
  height: 100%;
  min-height: 100vh;
  background: hsl(0 0% 100%);
  border: 0;
  border-radius: 0;
  box-shadow: none;
  overflow: hidden;
  transition: background 0.2s ease;
}

.sidebar-inner-collapsed {
  border-radius: 0;

  .sidebar-menu-wrap {
    padding: 12px 8px;
  }
}

.sidebar-menu-wrap {
  flex: 1;
  overflow: hidden;
  overflow-y: auto;
  padding: 14px 10px 12px;
  scrollbar-width: none;

  &::-webkit-scrollbar {
    width: 0;
    height: 0;
  }
  &::-webkit-scrollbar-thumb {
    background: rgb(148 163 184 / 0.35);
    border-radius: 2px;
  }
}

.sidebar-footer {
  flex-shrink: 0;
  border-top: 1px solid hsl(214 32% 91%);
  padding: 10px 12px;
}

.sidebar-collapse-btn {
  display: flex;
  width: 100%;
  height: 30px;
  align-items: center;
  justify-content: center;
  margin-top: 0;
  border: 0;
  border-radius: 6px;
  background: transparent;
  cursor: pointer;
  color: hsl(215 16% 47%);
  transition: background 0.15s ease, color 0.15s ease;

  svg {
    flex-shrink: 0;
    transition: transform 0.2s ease;
  }

  &:hover {
    background: hsl(210 40% 96%);
    color: hsl(222 47% 11%);
  }
}

.layout-content-main {
  margin: 0 20px 20px;
  position: relative;
  padding-top: 56px;
}

.layout-content-main-fix {
  padding-top: 56px;
}

.fluid-header {
  padding-top: 0;
}

.main-view-fix {
  padding-top: 50px;
}

.noMultiTabs {
  padding-top: 0;
}
</style>
