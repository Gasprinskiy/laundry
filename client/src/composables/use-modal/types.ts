import type { Component } from 'vue';

export interface UseModalState {
  component: Component | null;
  props?: any;
  emits?: any;
}

// export interface ShowModalOptions<T extends Component> {
//   component: T,
//   props?: InstanceType<T>["$props"],
//   emits?: Parameters<InstanceType<T>["$emit"]>
// }
