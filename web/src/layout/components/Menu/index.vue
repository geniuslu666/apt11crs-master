<template>
  <NMenu class="app-side-menu" :options="menus" :inverted="false" :mode="mode" :collapsed="collapsed" :collapsed-width="48"
    :collapsed-icon-size="18" :indent="26" :expanded-keys="openKeys" :value="getSelectedKeys"
    @update:value="clickMenuItem" @update:expanded-keys="menuExpanded" />
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, reactive, computed, watch, toRefs, PropType } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useAsyncRouteStore } from '@/store/modules/asyncRoute';
import { generatorMenu } from '@/utils';
import { useProjectSettingStore } from '@/store/modules/projectSetting';

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
      return selectedKeys.value;
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
      menus.value = generatorMenu(asyncRouteStore.getMenus);
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
      state.openKeys = openKeys;
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
  --nav-foreground: #192b47;
  --nav-muted: #71829d;
  --nav-subtle: #f6f9fc;
  --nav-active: #f3f7fb;
  --nav-border: #e5edf5;
  --nav-primary: #0f8fc9;
  --nav-radius: 8px;
  padding: 4px 0;

  :deep(.n-menu-item),
  :deep(.n-submenu) {
    margin: 2px 0;
  }

  :deep(.n-menu-item-content),
  :deep(.n-submenu > .n-menu-item-content) {
    height: 42px;
    padding-right: 14px !important;
    border-radius: var(--nav-radius);
    color: var(--nav-foreground);
    font-size: 16px;
    font-weight: 500;
    letter-spacing: 0;
    transition: background 0.15s ease, color 0.15s ease, box-shadow 0.15s ease;
  }

  :deep(.n-menu-item-content::before),
  :deep(.n-submenu > .n-menu-item-content::before) {
    display: none !important;
  }

  :deep(.n-menu-item-content:hover),
  :deep(.n-submenu > .n-menu-item-content:hover) {
    background: var(--nav-subtle);
    color: var(--nav-foreground);
  }

  :deep(.n-menu-item-content--selected),
  :deep(.n-menu-item-content--selected:hover),
  :deep(.n-menu-item-content--child-active) {
    background: var(--nav-active) !important;
    color: var(--nav-foreground) !important;
  }

  :deep(.n-menu-item-content--selected .n-menu-item-content-header),
  :deep(.n-menu-item-content--selected .n-menu-item-content__icon),
  :deep(.n-menu-item-content--selected .n-menu-item-content__arrow),
  :deep(.n-menu-item-content--child-active .n-menu-item-content-header),
  :deep(.n-menu-item-content--child-active .n-menu-item-content__icon),
  :deep(.n-menu-item-content--child-active .n-menu-item-content__arrow) {
    color: var(--nav-foreground) !important;
  }

  :deep(.n-menu-item-content__icon) {
    margin-right: 10px;
    color: currentColor;
    opacity: 0.72;
    width: 22px;
  }

  :deep(.n-menu-item-content__arrow) {
    color: var(--nav-muted);
    font-size: 18px;
  }

  :deep(.n-menu-item-content-header) {
    color: currentColor;
    line-height: 1;
    max-width: 6em;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  :deep(.n-submenu-children) {
    margin-left: 0;
    padding-left: 0;
    border-left: 0;
  }

  :deep(.n-submenu-children .n-menu-item-content),
  :deep(.n-submenu-children .n-submenu > .n-menu-item-content) {
    height: 36px;
    border-radius: 6px;
    color: var(--nav-foreground);
    font-size: 15px;
    font-weight: 500;
  }

  :deep(.n-submenu-children .n-menu-item-content__icon),
  :deep(.n-submenu-children .n-submenu > .n-menu-item-content .n-menu-item-content__icon) {
    display: none;
  }

  :deep(.n-submenu-children .n-menu-item-content-header) {
    padding-left: 10px;
  }

  :deep(.n-submenu-children .n-submenu-children .n-menu-item-content) {
    height: 32px;
    color: var(--nav-muted);
    font-size: 14px;
  }

  :deep(.n-submenu-children .n-submenu-children .n-menu-item-content-header) {
    padding-left: 24px;
  }

  :deep(.n-submenu-children .n-menu-item-content--selected),
  :deep(.n-submenu-children .n-menu-item-content--selected:hover) {
    background: transparent !important;
    color: var(--nav-primary) !important;
  }

  :deep(.n-submenu-children .n-menu-item-content--selected .n-menu-item-content-header) {
    color: var(--nav-primary) !important;
    font-weight: 600;
  }

  &.n-menu--collapsed :deep(.n-menu-item),
  &.n-menu--collapsed :deep(.n-submenu) {
    display: flex;
    justify-content: center;
  }

  &.n-menu--collapsed :deep(.n-menu-item-content),
  &.n-menu--collapsed :deep(.n-submenu > .n-menu-item-content) {
    width: 48px;
    height: 42px;
    padding: 0 !important;
    justify-content: center;
    border-radius: 8px;
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
