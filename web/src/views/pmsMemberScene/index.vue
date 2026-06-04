<template>
  <div class="member-admin-page">
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
          <n-tab-pane name="" tab="积分全局规则">
            <GlobalRule />
          </n-tab-pane>
          <n-tab-pane name="1" tab="民宿预定场景">
            <SceneRule sceneId="1"/>
          </n-tab-pane>
          <n-tab-pane name="2" tab="餐厅预定场景">
            <SceneRule sceneId="2"/>
          </n-tab-pane>
          <n-tab-pane name="3" tab="按摩预定场景">
            <SceneRule sceneId="3"/>
          </n-tab-pane>
          <n-tab-pane name="4" tab="接送机/包车预定场景">
            <SceneRule sceneId="4"/>
          </n-tab-pane>
          <n-tab-pane name="5" tab="储物柜预定场景">
            <SceneRule sceneId="5"/>
          </n-tab-pane>
        </n-tabs>

      </n-card>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import { useMessage } from 'naive-ui';
import {options, loadOptions} from "@/views/pmsMemberScene/model";
import GlobalRule from '@/views/pmsMemberScene/global_rule.vue';
import SceneRule from "@/views/pmsMemberScene/scene_rule.vue";


const group = ref('memberlevel');
const show = ref(false);
const rules = {};
const message = useMessage();





onMounted(() => {
  loadOptions();
  load();
});

function load() {

}
</script>
