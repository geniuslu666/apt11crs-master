<template>
  <div class="logo-wrap" :class="{ 'logo-collapsed': collapsed }">
    <span class="logo-mark" aria-hidden="true">
      <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M12 14v-4" />
        <path d="M9.2 12.4 7.8 11" />
        <path d="m14.8 12.4 1.4-1.4" />
        <path d="M19.4 14a7.8 7.8 0 1 0-14.8 0" />
        <path d="M5 18h14" />
      </svg>
    </span>
    <transition name="logo-text-fade">
      <div v-if="!collapsed" class="logo-text">
        <span class="logo-title">住一 CRS</span>
        <span class="logo-sub" :title="state.pmsmeun && propertyName ? propertyName : 'CRS 管理后台'">
          {{ state.pmsmeun && propertyName ? propertyName : 'CRS 管理后台' }}
        </span>
      </div>
    </transition>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref, watch, onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import { useUserStore } from '@/store/modules/user';
import { List } from '@/api/ten';
import { storage } from '@/utils/Storage';
import { getlang } from '@/utils/smjcomm';

export interface Props {
  collapsed?: boolean;
}
const props = withDefaults(defineProps<Props>(), { collapsed: false });

const userStore = useUserStore();
const currentRoute = useRoute();

const state = reactive({
  pmsmeun: false,
  chooseitem: userStore.getuserPms,
});

const propertyName = ref(
  userStore.getuserPms && userStore.getuserPms.id > 0
    ? namegetlang(userStore.getuserPms.nameLanguage)
    : ''
);

function namegetlang(data: any) {
  if (data) return getlang(data, userStore.language).content;
  return '';
}

watch(() => userStore.getuserPms, () => {
  propertyName.value = (userStore.getuserPms && userStore.getuserPms.id > 0)
    ? namegetlang(userStore.getuserPms.nameLanguage)
    : '';
}, { immediate: true, deep: true });

watch(() => currentRoute.fullPath, () => {
  state.pmsmeun = currentRoute.path.indexOf('pms/') !== -1;
});

const tenListLoad = async () => {
  const res = await List({ page: 1, pageSize: 200 });
  const wylists = res.list;
  storage.set('wylists', wylists);
  if (!userStore.getuserPms) {
    chooseItem(res.list[0]);
  }
};

const chooseItem = (item: any) => {
  state.chooseitem = item;
  storage.set('userPms', item);
  userStore.setuserPms(item);
};

onMounted(async () => {
  await tenListLoad();
  state.pmsmeun = currentRoute.path.indexOf('pms/') !== -1;
});
</script>

<style lang="less" scoped>
.logo-wrap {
  display: flex;
  align-items: center;
  min-height: 56px;
  gap: 12px;
  padding: 0 12px;
  overflow: hidden;
  flex-shrink: 0;
  border-bottom: 1px solid hsl(214 32% 91%);
  transition: padding 0.2s ease;

  &.logo-collapsed {
    padding: 0;
    justify-content: center;
    min-height: 56px;
    gap: 0;
  }
}

.logo-mark {
  display: inline-flex;
  width: 36px;
  height: 36px;
  align-items: center;
  justify-content: center;
  border-radius: 6px;
  background: #38aeea;
  color: #fff;
  flex-shrink: 0;

  svg {
    width: 20px;
    height: 20px;
  }
}

.logo-text {
  min-width: 0;
  overflow: hidden;
}

.logo-title {
  display: block;
  font-size: 14px;
  font-weight: 700;
  color: hsl(222 47% 11%);
  letter-spacing: 0;
  white-space: nowrap;
  line-height: 1.25;
}

.logo-sub {
  display: block;
  font-size: 12px;
  color: hsl(215 16% 47%);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  max-width: 160px;
  margin-top: 2px;
}

// Fade transition for logo text
.logo-text-fade-enter-active,
.logo-text-fade-leave-active {
  transition: opacity 0.15s ease;
}
.logo-text-fade-enter-from,
.logo-text-fade-leave-to {
  opacity: 0;
}
</style>
