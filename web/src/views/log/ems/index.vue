<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">邮件发送记录</text>
        </template>
        在这里，您可以方便地查看平台的所有邮箱发送记录
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{
                    padding: '0 20px 20px',
                  }">
      <BasicForm
        @register="register"
        @submit="handleSubmit"
        @reset="handleReset"
        ref="searchFormRef"
      >
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
      </BasicForm>

      <BasicTable
        :openChecked="false"
        :columns="columns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        ref="actionRef"
        @update:checked-row-keys="onCheckedRow"
        :scroll-x="scrollX"
        :resizeHeightOffset="-20000"
      >
        <template #tableTitle>
<!--          <n-button type="error" @click="batchDelete" :disabled="batchDeleteDisabled">
            <template #icon>
              <n-icon>
                <DeleteOutlined />
              </n-icon>
            </template>
            批量删除
          </n-button>-->
        </template>
      </BasicTable>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { computed, h, reactive, ref } from 'vue';
  import {NEllipsis, NTag, useDialog, useMessage} from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, FormSchema, useForm } from '@/components/Form/index';
  import { getLogList, Delete } from '@/api/log/emslog';
  import { DeleteOutlined } from '@vicons/antd';
  import { Dicts } from '@/api/dict/dict';
  import { adaTableScrollX, getOptionLabel, getOptionTag, Options } from '@/utils/hotgo';
  import { defRangeShortcuts } from '@/utils/dateUtil';

  const options = ref<Options>({
    config_sms_template: [],
  });

  const columns = [
    {
      title: 'ID',
      key: 'id',
      width: 100,
    },
    {
      title: '事件模板',
      key: 'event',
      render(row) {
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: getOptionTag(options.value.config_sms_template, row.event),
            bordered: false,
          },
          {
            default: () => getOptionLabel(options.value.config_sms_template, row.event),
          }
        );
      },
      width: 130,
    },
    {
      title: '邮箱',
      key: 'email',
      render(row) {
        return row.email;
      },
      width: 180,
    },
    {
      title: '验证码',
      key: 'code',
      width: 120,
    },
    {
      title: '验证次数',
      key: 'times',
      width: 100,
    },
    {
      title: '邮件内容',
      key: 'content',
      width: 300,
    },
    {
      title: '发送者IP',
      key: 'ip',
      width: 200,
    },
    {
      title: '状态',
      key: 'status',
      render(row) {
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: row.status == 2 ? 'success' : 'warning',
            bordered: false,
          },
          {
            default: () => (row.status == 2 ? '已验证' : '未验证'),
          }
        );
      },
      width: 100,
    },
    {
      title: '发送时间',
      key: 'createdAt',
      width: 160,
    },
    {
      title: '更新时间',
      key: 'updatedAt',
      width: 160,
    },
  ];

  const dialog = useDialog();
  const message = useMessage();
  const actionRef = ref();
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);
  const searchFormRef = ref<any>({});

  const schemas = ref<FormSchema[]>([
    /*{
      field: 'event',
      component: 'NSelect',
      label: '事件模板',
      componentProps: {
        placeholder: '请选择事件模板',
        options: [],
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },*/
    {
      field: 'email',
      component: 'NInput',
      label: '邮箱',
      componentProps: {
        placeholder: '请输入邮箱',
        onInput: (e: any) => {
          console.log(e);
        },
      },
      rules: [{ trigger: ['blur'] }],
    },
    {
      field: 'status',
      component: 'NSelect',
      label: '状态',
      componentProps: {
        placeholder: '请选择状态',
        options: [
          {
            label: '未验证',
            value: '1',
          },
          {
            label: '已验证',
            value: '2',
          },
        ],
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'createdAt',
      component: 'NDatePicker',
      label: '发送时间',
      componentProps: {
        type: 'datetimerange',
        clearable: true,
        shortcuts: defRangeShortcuts(),
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
  ]);

  const actionColumn = reactive({
    width: 80,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
          },
        ],
      });
    },
  });

  const scrollX = computed(() => {
    // return adaTableScrollX(columns, actionColumn.width);
    return adaTableScrollX(columns);
  });

  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
    labelWidth: 80,
    schemas,
  });

  function onCheckedRow(rowKeys) {
    batchDeleteDisabled.value = rowKeys.length <= 0;
    checkedIds.value = rowKeys;
  }

  function handleDelete(record: Recordable) {
    console.log('点击了删除', record);
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '不确定',
      onPositiveClick: () => {
        Delete(record).then((_res) => {
          message.success('操作成功');
          reloadTable();
        });
      },
      onNegativeClick: () => {
        // message.error('不确定');
      },
    });
  }

  function batchDelete() {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '不确定',
      onPositiveClick: () => {
        Delete({ id: checkedIds.value }).then((_res) => {
          message.success('操作成功');
          reloadTable();
        });
      },
      onNegativeClick: () => {
        // message.error('不确定');
      },
    });
  }

  const loadDataTable = async (res) => {
    await loadOptions();
    return await getLogList({ ...searchFormRef.value?.formModel, ...res });
  };

  function reloadTable() {
    actionRef.value.reload();
  }

  function handleSubmit(values: Recordable) {
    console.log(values);
    reloadTable();
  }

  function handleReset(values: Recordable) {
    console.log(values);
    reloadTable();
  }

  async function loadOptions() {
    options.value = await Dicts({
      types: ['config_sms_template'],
    });
    for (const item of schemas.value) {
      switch (item.field) {
        case 'event':
          item.componentProps.options = options.value.config_sms_template;
          break;
      }
    }
  }
</script>

<style lang="less" scoped></style>
