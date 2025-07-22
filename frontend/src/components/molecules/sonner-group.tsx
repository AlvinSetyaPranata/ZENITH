import {
  createEffect,
  createSignal,
  For,
  Index,
  onMount,
  Setter,
  Show,
} from "solid-js";
import Sooner from "../atoms/sooner";
import { Presence } from "solid-motionone";

type SoonerDataType = {
  id: number;
  title: string;
  desc: string;
  isRemoved: boolean;
};

interface SoonerGroupPropsType {
  soonersData: () => SoonerDataType[];
  soonersSetter: Setter<SoonerDataType[]>;
}

export default function SonnerGroup(props: SoonerGroupPropsType) {
  const onSonnerFinished = (id: number) => {
    if (props.soonersData().length == 1) {
      setTimeout(() => {
        props.soonersSetter((prevs) =>
          prevs.map((item) =>
            item.id === id ? { ...item, isRemoved: true } : item
          )
        );
      }, 900);

      return;
    }

    props.soonersSetter((prevs) =>
      prevs.map((item) =>
        item.id === id ? { ...item, isRemoved: true } : item
      )
    );
  };

  // createEffect(() => console.log(props.soonersData()), [props.soonersData]);

  return (
    <div class="fixed top-0 left-1/2 -translate-x-1/2 overflow-visible flex flex-col gap-y-8 pt-8 h-max">
      <For
        each={props.soonersData()}
        children={(item) => (
          <Presence>
            <Show keyed when={!item.isRemoved}>
              <Sooner
                id={item.id}
                title={item.title}
                desc={item.desc}
                setIsClosed={onSonnerFinished}
              />
            </Show>
          </Presence>
        )}
      />
    </div>
  );
}
