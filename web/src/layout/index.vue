<template>
  <n-layout class="layout" :position="fixedMenu" has-sider>
    <n-layout-sider v-if="
      !isMobile && isMixMenuNoneSub && (navMode === 'vertical' || navMode === 'horizontal-mix')
    " :position="fixedMenu"
      :collapsed="collapsed" collapse-mode="width" :collapsed-width="64" :width="leftMenuWidth"
      :native-scrollbar="false" :inverted="inverted" class="layout-sider">
      <div class="sidebar-inner">
        <Logo :collapsed="collapsed" />
        <div class="sidebar-menu-wrap">
          <AsideMenu v-model:collapsed="collapsed" v-model:location="getMenuLocation" />
        </div>
        <!-- Collapse toggle button -->
        <div class="sidebar-collapse-btn" @click="collapsed = !collapsed" :title="collapsed ? '展开菜单' : '折叠菜单'">
          <svg v-if="collapsed" xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m9 18 6-6-6-6"/></svg>
          <svg v-else xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m15 18-6-6 6-6"/></svg>
          <span v-if="!collapsed" class="collapse-label">折叠菜单</span>
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
  getNavMode,
  getNavTheme,
  getHeaderSetting,
  getMenuSetting,
  getMultiTabsSetting,
} = useProjectSetting();

const route = useRoute();
const settingStore = useProjectSettingStore();

const navMode = getNavMode;

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
  const { minMenuWidth, menuWidth } = unref(getMenuSetting);
  return collapsed.value ? minMenuWidth : menuWidth;
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
  background-color: #0a2540 !important;

  :deep(.n-drawer-body-content-wrapper) {
    padding: 0;
    background-color: #0a2540;
  }
}
</style>
<style lang="less" scoped>
.layout {
  display: flex;
  flex-direction: row;
  flex: auto;

  &-default-background {
    background: #f6f9fc;
  }

  .layout-sider {
    min-height: 100vh;
    box-shadow: none;
    border-right: none;
    position: relative;
    z-index: 13;
    transition: width 0.2s ease;
  }

  .layout-content {
    flex: auto;
    min-height: 100vh;
    background: #f6f9fc;
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
  background: #0a2540;
}

.sidebar-menu-wrap {
  flex: 1;
  overflow: hidden;
  overflow-y: auto;
  padding: 4px 0;

  &::-webkit-scrollbar {
    width: 4px;
  }
  &::-webkit-scrollbar-thumb {
    background: rgba(255,255,255,0.1);
    border-radius: 2px;
  }
}

.sidebar-collapse-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 16px;
  margin: 4px 8px 8px;
  border-radius: 6px;
  cursor: pointer;
  color: #697e99;
  font-size: 13px;
  transition: background 0.15s, color 0.15s;
  border-top: 1px solid rgba(255,255,255,0.06);
  padding-top: 14px;

  svg {
    flex-shrink: 0;
    transition: transform 0.2s ease;
  }

  .collapse-label {
    font-size: 13px;
    white-space: nowrap;
    overflow: hidden;
  }

  &:hover {
    background: rgba(255,255,255,0.06);
    color: #c8d2e0;
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
