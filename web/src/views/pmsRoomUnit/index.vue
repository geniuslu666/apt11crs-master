<template>
  <div>
    <div class="n-layout-page-header" style="margin: 0">
      <n-card :bordered="false"
              :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">房间</text>
        </template>
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard"
            :content-style="{
                    padding: '0 20px 20px',
                  }">
      <!-- <BasicForm
        ref="searchFormRef"
        @register="register"
        @submit="reloadTable"
        @reset="reloadTable"
        @keyup.enter="reloadTable"
      >
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
      </BasicForm> -->
      <BasicTable
        ref="actionRef"
        :openChecked="false"
        :columns="columns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        :actionColumn="actionColumn"
        :scroll-x="scrollX"
        :resizeHeightOffset="-10000"
        :checked-row-keys="checkedIds"
        @update:checked-row-keys="handleOnCheckedRow"
      >
      </BasicTable>
    </n-card>
    <Edit ref="editRef" @reloadTable="reloadTable" />
    <View ref="viewRef" />
  </div>
</template>

<script lang="ts" setup>
import {h, reactive, ref, computed, onMounted, watch} from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { usePermission } from '@/hooks/web/usePermission';
  import { List, Export, Delete } from '@/api/pmsRoomUnit';
  import { PlusOutlined, ExportOutlined, DeleteOutlined } from '@vicons/antd';
  import { columns, schemas } from './model';
  import { adaTableScrollX } from '@/utils/hotgo';
  import Edit from './edit.vue';
  import View from './view.vue';
  import { useRouter } from 'vue-router';
  import {useUserStore} from "@/store/modules/user";
  const userStore = useUserStore();
  const router = useRouter();
  const dialog = useDialog();
  const message = useMessage();
  const { hasPermission } = usePermission();
  const actionRef = ref();
  const searchFormRef = ref<any>({});
  const editRef = ref();
  const viewRef = ref();
  const checkedIds = ref([]);

  const actionColumn = reactive({
    rowdata: {},
    width: 216,
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
            auth: ['/pmsRoomUnit/view'],
          },
          {
            label: '编辑',
            onClick: handleEdit.bind(null, record),
            auth: ['/pmsRoomUnit/edit'],
          },

          // {
          //   label: '删除',
          //   onClick: handleDelete.bind(null, record),
          //   auth: ['/pmsRoomUnit/delete'],
          // },
        ],
        // dropDownActions: [
        //   {
        //     label: '查看详情',
        //     key: 'view',
        //     auth: ['/pmsRoomUnit/view'],
        //   },
        // ],
        // select: (key) => {
        //   if (key === 'view') {
        //     return handleView(record);
        //   }
        // },
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
    if (router.currentRoute.value.query?.uid) {
      res.puid = router.currentRoute.value.query.uid;
    }else{
      res.puid = userStore.getuserPms.uid;
    }
    return await List({ ...searchFormRef.value?.formModel, ...res });
  };

  // 更新选中的行
  function handleOnCheckedRow(rowKeys) {
    checkedIds.value = rowKeys;
  }

  // 重新加载表格数据
  function reloadTable() {
    actionRef.value?.reload();
  }

  watch(
    () => userStore.getuserPms,
    () => {
      if(!router.currentRoute.value.query.uid){
        reloadTable();
      }
    },
    {
      immediate: true,
      deep: true,
    }
  );

  // 添加数据
  function addTable() {
    editRef.value.openModal(null);
  }

  // 编辑数据
  function handleEdit(record: Recordable) {
    editRef.value.openModal(record);
  }

  // 查看详情
  function handleView(record: Recordable) {
    viewRef.value.openModal(record);
  }

  // 单个删除
  function handleDelete(record: Recordable) {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete(record).then((_res) => {
          message.success('删除成功');
          reloadTable();
        });
      },
    });
  }

  // 批量删除
  function handleBatchDelete() {
    if (checkedIds.value.length < 1) {
      message.error('请至少选择一项要删除的数据');
      return;
    }

    dialog.warning({
      title: '警告',
      content: '你确定要批量删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete({ id: checkedIds.value }).then((_res) => {
          checkedIds.value = [];
          message.success('删除成功');
          reloadTable();
        });
      },
    });
  }

  // 导出
  function handleExport() {
    message.loading('正在导出列表...', { duration: 1200 });
    Export(searchFormRef.value?.formModel);
  }
  onMounted(() => {
    if (router.currentRoute.value.query?.rowdata) {
      actionColumn.rowdata = JSON.parse(router.currentRoute.value.query.rowdata);
    }
  });
</script>

<style lang="less" scoped></style>
