<template>
  <div class="tableAction">
    <div class="flex items-center justify-center">
      <template v-for="(action, index) in getActions" :key="`${index}-${action.label}`">
        <n-button v-bind="action" class="mx-1">
          {{ action.label }}
          <template #icon v-if="action.hasOwnProperty('icon')">
            <n-icon :component="action.icon" />
          </template>
        </n-button>
      </template>
      <n-dropdown v-if="dropDownActions && getDropdownList.length" trigger="hover" :options="getDropdownList"
        @select="select">
        <slot name="more"></slot>
        <n-button v-bind="getMoreProps" class="mx-1" v-if="!$slots.more" icon-placement="right">
          <div class="flex items-center">
            <span>更多</span>
            <n-icon size="14" class="ml-1">
              <DownOutlined />
            </n-icon>
          </div>
          <!--          <template #icon> </template>-->
        </n-button>
      </n-dropdown>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, PropType, toRaw, ref, onMounted, onUnmounted } from 'vue';
import { ActionItem } from '@/components/Table';
import { usePermission } from '@/hooks/web/usePermission';
import { isBoolean, isFunction } from '@/utils/is';
import { DownOutlined } from '@vicons/antd';

export default defineComponent({
  name: 'TableAction',
  components: { DownOutlined },
  props: {
    actions: {
      type: Array as PropType<ActionItem[]>,
      default: null,
      required: true,
    },
    dropDownActions: {
      type: Array as PropType<ActionItem[]>,
      default: null,
    },
    style: {
      type: String as PropType<String>,
      default: 'button',
    },
    select: {
      type: Function as PropType<Function>,
      default: () => { },
    },
    size: {
      type: String as PropType<string>,
      default: undefined,
    },
  },
  setup(props) {
    const { hasPermission } = usePermission();

    // 检测是否为移动端（<768px）
    const isMobile = ref(window.innerWidth < 768);

    const handleResize = () => {
      isMobile.value = window.innerWidth < 768;
    };

    onMounted(() => {
      window.addEventListener('resize', handleResize);
    });

    onUnmounted(() => {
      window.removeEventListener('resize', handleResize);
    });

    // 获取默认size：移动端small，PC端medium；如果props.size有值则优先使用props.size
    const defaultSize = computed(() => {
      if (props.size) {
        return props.size;
      }
      return isMobile.value ? 'small' : 'medium';
    });

    const actionType =
      props.style === 'button' ? 'default' : props.style === 'text' ? 'primary' : 'default';
    const actionText =
      props.style === 'button' ? undefined : props.style === 'text' ? true : undefined;

    const getMoreProps = computed(() => {
      return {
        text: actionText,
        type: actionType,
        size: defaultSize.value,
      };
    });

    const getDropdownList = computed(() => {
      return (toRaw(props.dropDownActions) || [])
        .filter((action) => {
          return hasPermission(action.auth as string[]) && isIfShow(action);
        })
        .map((action) => {
          const { popConfirm, size: actionSize, ...restAction } = action;
          const finalSize = actionSize || defaultSize.value;
          return {
            ...restAction,
            text: actionText,
            type: getBtnType(action), //actionType,
            ...popConfirm,
            // 确保size属性最后设置，不会被action中的属性覆盖
            size: finalSize,
            onConfirm: popConfirm?.confirm,
            onCancel: popConfirm?.cancel,
          };
        });
    });

    function isIfShow(action: ActionItem): boolean {
      const ifShow = action.ifShow;

      let isIfShow = true;

      if (isBoolean(ifShow)) {
        isIfShow = ifShow;
      }
      if (isFunction(ifShow)) {
        isIfShow = ifShow(action);
      }
      return isIfShow;
    }

    function getBtnType(action) {
      if (action.type !== undefined && action.type !== '') {
        return action.type;
      }
      switch (action.label) {
        case '编辑':
          return 'success';
        case '启用':
        case '已禁用':
          return 'success';
        case '已启用':
        case '禁用':
          return 'warning';
        case '删除':
          return 'error';
        case '查看详情':
          return 'default';
        default:
          return 'primary';
      }
    }

    function getBtnColor(action) {
      if (action.color !== undefined && action.color !== '') {
        return action.color;
      }
      return '';
    }

    const getActions = computed(() => {
      return (toRaw(props.actions) || [])
        .filter((action) => {
          return hasPermission(action.auth as string[]) && isIfShow(action);
        })
        .map((action) => {
          const { popConfirm, size: actionSize, ...restAction } = action;
          //需要展示什么风格，自己修改一下参数
          // 在移动端强制使用small，除非action中明确指定了size
          const finalSize = actionSize || defaultSize.value;
          return {
            ...restAction,
            text: actionText,
            type: getBtnType(action), //actionType,
            color: getBtnColor(action), //actionType,
            ...(popConfirm || {}),
            // 确保size属性最后设置，不会被action中的属性覆盖
            size: finalSize,
            onConfirm: popConfirm?.confirm,
            onCancel: popConfirm?.cancel,
            enable: !!popConfirm,
          };
        });
    });

    return {
      getActions,
      getDropdownList,
      getMoreProps,
    };
  },
});
</script>
