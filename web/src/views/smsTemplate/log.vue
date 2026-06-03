<template>
  <div>
    <div class="n-layout-page-header" style="margin: 0">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">发送记录</text>
        </template>
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{
                    padding: '0 20px 20px',
                  }">
      <BasicForm  ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable">
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
      </BasicForm>
      <BasicTable  ref="actionRef" :openChecked="false" :columns="columns" :request="loadDataTable" :scroll-x="scrollX" :resizeHeightOffset="-10000">
      </BasicTable>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted } from 'vue';
import { BasicTable } from '@/components/Table';
import { Log } from '@/api/smsTemplate';
import { columns, schemas } from './log_model';
import { adaTableScrollX } from '@/utils/hotgo';
import {BasicForm, useForm} from "@/components/Form";

const actionRef = ref();
const searchFormRef = ref<any>({});

const [register, {}] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas,
});

const scrollX = computed(() => {
  return adaTableScrollX(columns, 0);
});

// 加载表格数据
const loadDataTable = async (res) => {
  return await Log({ ...searchFormRef.value?.formModel, ...res });
};

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

onMounted(() => {
});
</script>

<style lang="less" scoped></style>

