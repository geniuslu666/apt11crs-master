<template>
  <NMenu class="app-side-menu" :options="menus" :inverted="false" :mode="mode" :collapsed="collapsed" :collapsed-width="40"
    :collapsed-icon-size="16" :indent="16" :expanded-keys="openKeys" :value="getSelectedKeys"
    @update:value="clickMenuItem" @update:expanded-keys="menuExpanded" />
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, reactive, computed, watch, toRefs, unref, PropType } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useAsyncRouteStore } from '@/store/modules/asyncRoute';
import { generatorMenu, generatorMenuMix } from '@/utils';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { useProjectSetting } from '@/hooks/setting/useProjectSetting';

export default defineComponent({
  name: 'Menu',
  components: {},
  props: {
    mode: {
      // 菜单模式
      type: String,
      default: 'vertical',
    },
    collapsed: {
      // 侧边栏菜单是否收起
      type: Boolean,
    },
    //位置
    location: {
      type: String,
      default: 'left',
    },
    // 可选：强制指定菜单主题（不传则跟随系统 navTheme）
    inverted: {
      type: Boolean as PropType<boolean>,
      default: undefined,
    },
  },
  emits: ['update:collapsed', 'clickMenuItem'],
  setup(props, { emit }) {
    // 当前路由
    const currentRoute = useRoute();
    const router = useRouter();
    const asyncRouteStore = useAsyncRouteStore();
    const settingStore = useProjectSettingStore();
    const menus = ref<any[]>([]);
    const selectedKeys = ref<string>(currentRoute.name as string);
    const headerMenuSelectKey = ref<string>('');

    const { getNavMode } = useProjectSetting();

    const navMode = getNavMode;

    // 获取当前打开的子菜单
    const matched = currentRoute.matched;

    const getOpenKeys = matched && matched.length ? matched.map((item) => item.name) : [];

    const state = reactive({
      openKeys: getOpenKeys,
    });

    // 自动根据全局主题计算的反色模式
    const autoInverted = computed(() => {
      return ['dark', 'header-dark'].includes(settingStore.navTheme);
    });

    // 最终给 NMenu 使用的 inverted，优先使用 props.inverted
    const menuInverted = computed(() => {
      return props.inverted ?? autoInverted.value;
    });

    const getSelectedKeys = computed(() => {
      let location = props.location;
      return location === 'left' || (location === 'header' && unref(navMode) === 'horizontal')
        ? unref(selectedKeys)
        : unref(headerMenuSelectKey);
    });

    // 监听分割菜单
    watch(
      () => settingStore.menuSetting.mixMenu,
      () => {
        updateMenu();
        if (props.collapsed) {
          emit('update:collapsed', !props.collapsed);
        }
      }
    );

    // 监听菜单收缩状态
    // watch(
    //   () => props.collapsed,
    //   (newVal) => {
    //   }
    // );

    // 跟随页面路由变化，切换菜单选中状态
    watch(
      () => currentRoute.fullPath,
      () => {
        updateMenu();
        const matched = currentRoute.matched;
        state.openKeys = matched.map((item) => item.name);
        const activeMenu: string = (currentRoute.meta?.activeMenu as string) || '';
        selectedKeys.value = activeMenu ? (activeMenu as string) : (currentRoute.name as string);
      }
    );

    function updateMenu() {
      if (!settingStore.menuSetting.mixMenu) {
        menus.value = generatorMenu(asyncRouteStore.getMenus);
      } else {
        //混合菜单
        const firstRouteName: string = (currentRoute.matched[0].name as string) || '';
        menus.value = generatorMenuMix(asyncRouteStore.getMenus, firstRouteName, props.location);
        const activeMenu: string = currentRoute?.matched[0].meta?.activeMenu as string;
        headerMenuSelectKey.value = (activeMenu ? activeMenu : firstRouteName) || '';
        ;
      }
    }

    // 点击菜单
    function clickMenuItem(key: string) {
      if (/http(s)?:/.test(key)) {
        window.open(key);
      } else {
        if (getSelectedKeys.value !== key) {
          router.push({ name: key });
        }
      }
      emit('clickMenuItem' as any, key);
    }

    //展开菜单
    function menuExpanded(openKeys: string[]) {
      if (!openKeys) return;
      const latestOpenKey = openKeys.find((key) => state.openKeys.indexOf(key) === -1);
      const isExistChildren = findChildrenLen(latestOpenKey as string);
      state.openKeys = isExistChildren ? (latestOpenKey ? [latestOpenKey] : []) : openKeys;
    }

    //查找是否存在子路由
    function findChildrenLen(key: string) {
      if (!key) return false;
      const subRouteChildren: string[] = [];
      for (const { children, key } of unref(menus)) {
        if (children && children.length) {
          subRouteChildren.push(key as string);
        }
      }
      return subRouteChildren.includes(key);
    }

    onMounted(() => {
      updateMenu();
      const matched = currentRoute.matched;
      state.openKeys = matched.map((item) => item.name);
      const activeMenu: string = (currentRoute.meta?.activeMenu as string) || '';
      selectedKeys.value = activeMenu ? (activeMenu as string) : (currentRoute.name as string);
    });

    return {
      ...toRefs(state),
      menuInverted,
      menus,
      selectedKeys,
      headerMenuSelectKey,
      getSelectedKeys,
      clickMenuItem,
      menuExpanded,
    };
  },
});
</script>

<style lang="less" scoped>
.app-side-menu {
  background: transparent;
  --smartcam-foreground: #152033;
  --smartcam-muted: #7b93ad;
  --smartcam-muted-bg: #f4f8fb;
  --smartcam-active-bg: #eaf7ff;
  --smartcam-border: #dce7f2;
  --smartcam-primary: #128fc8;

  :deep(.n-menu-item),
  :deep(.n-submenu) {
    margin: 2px 0;
  }

  :deep(.n-menu-item-content),
  :deep(.n-submenu > .n-menu-item-content) {
    height: 36px;
    padding-right: 8px !important;
    border-radius: 6px;
    color: var(--smartcam-muted);
    font-size: 14px;
    font-weight: 500;
    transition: background 0.15s ease, color 0.15s ease;
  }

  :deep(.n-menu-item-content::before),
  :deep(.n-submenu > .n-menu-item-content::before) {
    display: none !important;
  }

  :deep(.n-menu-item-content:hover),
  :deep(.n-submenu > .n-menu-item-content:hover) {
    background: var(--smartcam-muted-bg);
    color: var(--smartcam-foreground);
  }

  :deep(.n-menu-item-content--selected),
  :deep(.n-menu-item-content--selected:hover),
  :deep(.n-menu-item-content--child-active) {
    background: var(--smartcam-active-bg) !important;
    color: var(--smartcam-primary) !important;
  }

  :deep(.n-menu-item-content--selected .n-menu-item-content-header),
  :deep(.n-menu-item-content--selected .n-menu-item-content__icon),
  :deep(.n-menu-item-content--selected .n-menu-item-content__arrow),
  :deep(.n-menu-item-content--child-active .n-menu-item-content-header),
  :deep(.n-menu-item-content--child-active .n-menu-item-content__icon),
  :deep(.n-menu-item-content--child-active .n-menu-item-content__arrow) {
    color: var(--smartcam-primary) !important;
  }

  :deep(.n-menu-item-content__icon) {
    margin-right: 10px;
    color: currentColor;
  }

  :deep(.n-menu-item-content__arrow) {
    color: currentColor;
  }

  :deep(.n-menu-item-content-header) {
    color: currentColor;
    line-height: 1;
  }

  :deep(.n-submenu-children) {
    margin-left: 8px;
    padding-left: 8px;
    border-left: 1px solid var(--smartcam-border);
  }

  &.n-menu--collapsed :deep(.n-menu-item),
  &.n-menu--collapsed :deep(.n-submenu) {
    display: flex;
    justify-content: center;
  }

  &.n-menu--collapsed :deep(.n-menu-item-content),
  &.n-menu--collapsed :deep(.n-submenu > .n-menu-item-content) {
    width: 40px;
    height: 36px;
    padding: 0 !important;
    justify-content: center;
    border-radius: 6px;
  }

  &.n-menu--collapsed :deep(.n-menu-item-content__icon) {
    margin-right: 0;
  }

  &.n-menu--collapsed :deep(.n-submenu-children) {
    margin-left: 0;
    padding-left: 0;
    border-left: 0;
  }
}
</style>
