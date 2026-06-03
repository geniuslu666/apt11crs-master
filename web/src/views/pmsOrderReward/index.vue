<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card
        :bordered="false"
        class="proCard"
        size="small"
        :segmented="{ content: true }"
      >

        <n-tabs
          type="card"
          class="card-tabs"
          animated
        >
          <n-tab-pane name="1" tab="民宿预定场景">
            <Reward sceneId="1"/>
          </n-tab-pane>
        </n-tabs>

      </n-card>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import Reward from "@/views/pmsOrderReward/reward.vue";


const show = ref(false);

onMounted(() => {
});

function load() {

}
</script>
