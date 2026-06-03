<template>
  <div>
    <div class="n-layout-page-header" style="margin: 0">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">商户列表</text>
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
      <BasicTable  ref="actionRef" :openChecked="false" :columns="columns" :request="loadDataTable" :row-key="(row) => row.id" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000"  :checked-row-keys="checkedIds" @update:checked-row-keys="handleOnCheckedRow">
        <template #tableTitle>
          <n-button type="primary"  @click="addTable" class="min-left-space" v-if="hasPermission(['/thMch/edit'])">
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加商户
          </n-button>
        </template>
      </BasicTable>
    </n-card>
    <Edit ref="editRef" @reloadTable="reloadTable" />
  </div>
</template>

<script lang="ts" setup>
import { h, reactive, ref, computed, onMounted } from 'vue';
import { BasicTable, TableAction } from '@/components/Table';
import { usePermission } from '@/hooks/web/usePermission';
import {List, Switch} from '@/api/thMch';
import { PlusOutlined } from '@vicons/antd';
import {schemas, loadOptions, options} from './model';
import {adaTableScrollX, getOptionLabel, getOptionTag} from '@/utils/hotgo';
import Edit from './edit.vue';
import {BasicForm, useForm} from "@/components/Form";
import {useRouter} from "vue-router";
import {isNullObject} from "@/utils/is";
import {NButton, NSwitch, NTag, useMessage} from "naive-ui";

const message = useMessage();
const router = useRouter();
const { hasPermission } = usePermission();
const actionRef = ref();
const searchFormRef = ref<any>({});
const editRef = ref();
const checkedIds = ref([]);

const actionColumn = reactive({
  width: 216,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '商户详情',
          onClick: handleView.bind(null, record),
          auth: ['/thMch/view'],
        },
        {
          label: '编辑',
          onClick: handleEdit.bind(null, record),
          auth: ['/thMch/edit'],
        },
      ],
    });
  },
});

const columns = [
  {
    title: '商户名称',
    key: 'name',
    align: 'left',
    width: 155,
  },
  {
    title: '分类',
    key: 'thMchCategoryName',
    align: 'left',
    width: 155,
  },
  {
    title: '状态',
    key: 'status',
    align: 'left',
    width: 80,
    render(row) {
      if (isNullObject(row.status)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.sys_normal_disable, row.status),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.sys_normal_disable, row.status),
        }
      );
    },
  },

  {
    title: '门店数量',
    key: 'storeNum',
    align: 'left',
    width: 80,
    render(row){
      return h(
        NButton,
        {
          strong: true,
          size: 'small',
          text: true,
          type: 'primary',
          onClick: () => handleView({id:row.id}),
        },
        { default: () => row.storeOnNum + row.storeOffNum }
      )
    }
  },
  {
    title: '联系信息',
    key: 'contactInfo',
    align: 'left',
    width: 155,
    render(row){
      return h(
        'span',
        {
          style:{
            whiteSpace: "pre-line"
          }
        },
        {
          default: () => row.contactInfo
        }
      )
    }
  },
  {
    title: '是否启用',
    key: 'status',
    align: 'left',
    width: 80,
    render(row) {
      return h(NSwitch, {
        value: row.status === 1,
        checked: '开启',
        unchecked: '关闭',
        onUpdateValue: function (e) {
          row.status = e ? 1 : 2;
          Switch({ id: row.id, key: 'switch', status: row.status }).then((_res) => {
            message.success('操作成功');
          });
        },
      });
    },
  },
];

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

// 更新选中的行
function handleOnCheckedRow(rowKeys) {
  checkedIds.value = rowKeys;
}

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

// 添加数据
function addTable() {
  editRef.value.openModal(null);
}

// 查看详情
function handleView(record: Recordable) {
  router.push({ name: 'thMchView', params: { id: record.id } });
}

// 编辑数据
function handleEdit(record: Recordable) {
  editRef.value.openModal(record);
}

onMounted(() => {
  loadOptions();
});
</script>

<style lang="less" scoped></style>

