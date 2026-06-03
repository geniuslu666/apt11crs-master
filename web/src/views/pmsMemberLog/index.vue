<template>
  <div>
    <div class="n-layout-page-header" style="margin: 0">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">会员登录日志</text>
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
      <BasicTable  ref="actionRef"  :columns="columns" :request="loadDataTable" :row-key="(row) => row.id" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000"   >
        <template #tableTitle>
          <n-button type="primary" @click="handleExport" class="min-left-space" v-if="hasPermission(['/pmsMemberLog/export'])">
            <template #icon>
              <n-icon>
                <ExportOutlined />
              </n-icon>
            </template>
            导出
          </n-button>
        </template>
      </BasicTable>
    </n-card>
    <View ref="viewRef" />
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref, computed, onMounted } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { usePermission } from '@/hooks/web/usePermission';
  import { List, Export } from '@/api/pmsMemberLog';
  import { ExportOutlined } from '@vicons/antd';
  import { columns, schemas, loadOptions } from './model';
  import { adaTableScrollX } from '@/utils/hotgo';
  import View from './view.vue';

  const dialog = useDialog();
  const message = useMessage();
  const { hasPermission } = usePermission();
  const actionRef = ref();
  const searchFormRef = ref<any>({});
  const editRef = ref();
  const viewRef = ref();

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
            auth: ['/pmsMemberLog/view'],
          },
        ],
        /*dropDownActions: [
          {
            label: '详情',
            key: 'view',
            auth: ['/pmsMemberLog/view'],
          },
        ],
        select: (key) => {
          if (key === 'view') {
            return handleView(record);
          }
        },*/
      });
    },
  });

  const scrollX = computed(() => {
    return adaTableScrollX(columns, actionColumn.width);
  });

  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
    labelWidth: 80,
    schemas,
  });

  // 加载表格数据
  const loadDataTable = async (res) => {
    return await List({ ...searchFormRef.value?.formModel, ...res });
  };

  // 重新加载表格数据
  function reloadTable() {
    actionRef.value?.reload();
  }

  // 查看详情
  function handleView(record: Recordable) {
    viewRef.value.openModal(record);
  }

  // 导出
  function handleExport() {
    message.loading('正在导出列表...', { duration: 1200 });
    Export(searchFormRef.value?.formModel);
  }

  onMounted(() => {
    loadOptions();
  });
</script>

<style lang="less" scoped></style>
