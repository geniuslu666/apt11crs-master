<template>
  <div>
    <div class="n-layout-page-header" style="margin: 0">
      <n-card :bordered="false"
              :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">价格计划</text>
        </template>
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard"
            :content-style="{
                    padding: '0 20px 20px',
                  }">
      <BasicForm  ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable">
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
      </BasicForm>
      <div class="statusTab">
        <div :class="searchState == 0 ? 'active' : ''"><span @click="handlePlanState(0)">价格计划</span></div>
        <div :class="searchState == 1 ? 'active' : ''"><span @click="handlePlanState(1)">回收站</span></div>
      </div>
      <BasicTable  ref="actionRef" :openChecked="false" :columns="columns" :request="loadDataTable" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000">
        <template #tableTitle>
          <n-button type="primary"  @click="addTable" class="min-left-space" v-if="hasPermission(['/pricePlan/edit'])&&searchState==0" style="margin-top:5px">
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            新增计划
          </n-button>
        </template>
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
import { List, Status, Delete, Recycle } from '@/api/pmsPricePlan';
import { PlusOutlined } from '@vicons/antd';
import { columns, schemas, getRoomTypeList } from './model';
import { adaTableScrollX } from '@/utils/hotgo';
import Edit from "@/views/pmsPricePlan/edit.vue";
import View from "@/views/pmsPricePlan/view.vue";
import {useRouter} from "vue-router";
import {useUserStore} from "@/store/modules/user";

const router = useRouter();
const dialog = useDialog();
const message = useMessage();
const { hasPermission } = usePermission();
const actionRef = ref();
const searchFormRef = ref<any>({});
const editRef = ref();
const viewRef = ref();
const userStore = useUserStore();
const searchState = ref(0);

const actionColumn = reactive({
  width: 288,
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
          auth: ['/pricePlan/view'],
        },
        {
          label: '编辑',
          onClick: handleEdit.bind(null, record),
          ifShow: () => {
            return record.deletedAt  == null;
          },
          auth: ['/pricePlan/edit'],
        },
        {
          label: '禁用',
          onClick: handleStatus.bind(null, record, "N"),
          ifShow: () => {
            return record.pricePlanStatus === "Y" && record.deletedAt  == null;
          },
          auth: ['/pricePlan/status'],
        },
        {
          label: '启用',
          onClick: handleStatus.bind(null, record, "Y"),
          ifShow: () => {
            return record.pricePlanStatus === "N" && record.deletedAt  == null;
          },
          auth: ['/pricePlan/status'],
        },
        {
          label: '删除',
          onClick: handleDelete.bind(null, record),
          ifShow: () => {
            return record.deletedAt  == null;
          },
          auth: ['/pricePlan/delete'],
        },
        {
          label: '恢复',
          onClick: handleRecycle.bind(null, record),
          type: 'warning',
          auth: ['/pricePlan/recycle'],
          ifShow: () => {
            return record.deletedAt != null;
          },
        },
      ],
    });
  },
});

const scrollX = computed(() => {
  return adaTableScrollX(columns, actionColumn.width);
});

const [register, {}] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 120,
  schemas,
});

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

// 加载表格数据
const loadDataTable = async (res) => {
  if (router.currentRoute.value.query?.uid) {
    res.propertyId = router.currentRoute.value.query.uid;
  }else{
    res.propertyId = userStore.getuserPms.uid;
  }

  if(searchState.value > 0){
    res.type = searchState.value
  }

  return await List({ ...searchFormRef.value?.formModel, ...res });
};

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

// 添加数据
function addTable() {
  if (router.currentRoute.value.query?.uid) {
    var puid = router.currentRoute.value.query.uid;
  }else{
    var puid = userStore.getuserPms.uid;
  }
  editRef.value.openModal({propertyId: puid});
}

// 查看详情
function handleView(record: Recordable) {
  viewRef.value.openModal(record);
}

// 编辑数据
function handleEdit(record: Recordable) {
  if (router.currentRoute.value.query?.uid) {
    record.propertyId = router.currentRoute.value.query.uid;
  }else{
    record.propertyId = userStore.getuserPms.uid;
  }
  editRef.value.openModal(record);
}

// 修改状态
function handleStatus(record: Recordable, status: string) {
  Status({ id: record.id, status: status }).then((_res) => {
    message.success('修改成功');
    setTimeout(() => {
      reloadTable();
    });
  });
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

function handlePlanState(e){
  searchState.value = e;
  reloadTable()
}

// 单个恢复
function handleRecycle(record: Recordable) {
  dialog.warning({
    title: '警告',
    content: '你确定要恢复？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Recycle(record).then((_res) => {
        message.success('恢复成功');
        reloadTable();
      });
    },
  });
}

onMounted(() => {
  if (router.currentRoute.value.query?.uid) {
    var propertyId = router.currentRoute.value.query.uid;
  }else{
    var propertyId = userStore.getuserPms.uid;
  }
  getRoomTypeList(propertyId);
});
</script>

<style lang="less" scoped>
.statusTab{
  display: flex;
  margin-bottom: 4px;
  div{
    margin-right: 5px;
    padding: 0 16px;
    height: 30px;
    line-height: 30px;
    text-align: center;
    font-size: 14px;
    color: #4E5969;
    span{
      cursor: pointer;
    }
    &.active{
      background: #F2F3F8;
      border-radius: 30px;
      color: #1664FF;
      font-weight: 500;
    }
  }
}
</style>
