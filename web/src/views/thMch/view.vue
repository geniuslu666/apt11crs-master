<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-grid :y-gap="15" :cols="1">
        <n-gi>
          <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }" :content-style="{
                    padding: '0 20px 20px',
                  }">
            <template #header>
              <text style="font-weight: 500;font-size: 20px;color: #3D3D3D;line-height: 28px;">商户概要</text>
            </template>
            <n-descriptions label-placement="top" style="margin-top: 10px">
              <n-descriptions-item :span="3">
                商户名称： {{ data.name }}
              </n-descriptions-item>
              <n-descriptions-item :span="3">
                商户分类： {{ data.thMchCategoryName }}
              </n-descriptions-item>
              <n-descriptions-item :span="3">
                <div style="display: flex;align-items: center">
                  <span>商户LOGO：</span>
                  <img
                    :src="data.logo"
                    loading="lazy"
                    style="width: 80px;"
                  />
                </div>
              </n-descriptions-item>
              <n-descriptions-item :span="3">
                <div style="display: flex">
                  <span>联系信息： </span>
                  <div style="white-space: pre-line">{{ data.contactInfo }}</div>
                </div>
              </n-descriptions-item>
            </n-descriptions>
          </n-card>
        </n-gi>
        <n-gi>
          <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }" :content-style="{
                    padding: '0 20px 20px',
                  }">
            <template #header>
              <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">门店列表</text>
            </template>
            <BasicTable  ref="actionRef" :openChecked="false" :columns="columns" :request="loadDataTable" :row-key="(row) => row.id" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000">
              <template #tableTitle>
                <n-button type="primary"  @click="addTable" class="min-left-space" v-if="hasPermission(['/thMchStore/edit'])">
                  <template #icon>
                    <n-icon>
                      <PlusOutlined />
                    </n-icon>
                  </template>
                  添加门店
                </n-button>
              </template>
            </BasicTable>
          </n-card>
        </n-gi>
      </n-grid>
    </n-spin>

    <Edit ref="editRef" @reloadTable="reloadTable"/>
    <StoreView ref="viewRef" @reloadTable="reloadTable" />
  </div>
</template>

<script lang="ts" setup>
import {computed, h, onMounted, reactive, ref} from "vue";
import { View } from '@/api/thMch';
import {useRouter} from "vue-router";
import {useMessage} from "naive-ui";
import {columns, loadOptions} from "@/views/thMch/store_model";
import {PlusOutlined} from "@vicons/antd";
import {BasicTable, TableAction} from "@/components/Table";
import { List } from '@/api/thMchStore';
import {adaTableScrollX} from "@/utils/hotgo";
import Edit from "@/views/thMchStore/edit.vue";
import StoreView from "@/views/thMchStore/view.vue";
import {usePermission} from "@/hooks/web/usePermission";
import {loadOptions as loadTerminalOptions} from "@/views/terminal/model";

const { hasPermission } = usePermission();
const loading = ref(false);
const router = useRouter();
const message = useMessage();
const params = router.currentRoute.value.params;
const data = ref({});
const actionRef = ref();
const editRef = ref();
const viewRef = ref();

// 加载表格数据
const loadDataTable = async (res) => {
  res.mchId = params.id
  return await List({ ...res });
};

const actionColumn = reactive({
  width: 100,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '详情',
          onClick: handleView.bind(null, record),
          auth: ['/thMchStore/view'],
        },
      ],
    });
  },
});

const scrollX = computed(() => {
  return adaTableScrollX(columns, actionColumn.width);
});

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

// 添加数据
function addTable() {
  editRef.value.openModal(null,parseInt(params.id));
}

// 查看详情
function handleView(record: Recordable) {
  viewRef.value.openModal(record.id);
}

async function getInfo(){
  const res = await View(params);
  data.value = res;
}

onMounted(async () => {
  if (!params.id) {
    message.error('商户ID不正确，请检查！');
    return;
  }
  loadOptions();
  loadTerminalOptions()
  loading.value = true;
  await getInfo();
  loading.value = false;
});
</script>

<style lang="less" scoped>
</style>


