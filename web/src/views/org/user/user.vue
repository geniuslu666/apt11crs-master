<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">后台用户</text>
        </template>
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{
                    padding: '0 20px 20px',
                  }">
      <n-tabs
        type="card"
        class="card-tabs"
        :value="defaultTab"
        animated
        @before-leave="handleBeforeLeave"
      >
        <n-tab-pane
          :name="item.id.toString()"
          :tab="item.name"
          v-for="item in options.roleTabs"
          :key="item.key"
        >
          <List :type="defaultTab" />
        </n-tab-pane>
      </n-tabs>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, ref } from 'vue';
  import List from './list.vue';
  import { useRouter } from 'vue-router';
  import { options } from './model';

  const router = useRouter();
  const defaultTab = ref('-1');

  onMounted(() => {
    if (router.currentRoute.value.query?.type) {
      defaultTab.value = router.currentRoute.value.query.type as string;
    }
  });

  function handleBeforeLeave(tabName: string): boolean | Promise<boolean> {
    defaultTab.value = tabName;
    return true;
  }
</script>
