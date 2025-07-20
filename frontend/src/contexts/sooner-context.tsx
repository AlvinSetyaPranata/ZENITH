/* @refresh reload */

import {
  createContext,
  createEffect,
  createSignal,
  ParentProps,
  Show,
  useContext,
} from "solid-js";
import Sooner from "~/components/atoms/sooner";
import { AuthContext } from "./auth-context";
import { Presence } from "solid-motionone";
import SonnerGroup from "~/components/molecules/sonner-group";

interface SoonerType {
  id: number;
  title: string;
  desc: string;
  isRemoved: boolean
}

interface SoonerContextType {
  show: (title: string, desc: string) => void;
}

export const SoonerContext = createContext<SoonerContextType>();

export function SoonerProvider(props: ParentProps) {

  const [sooners, setSooners] = createSignal<SoonerType[]>([])

  const show = (title: string, desc: string) => {

    if (sooners().every((sooner) => sooner.isRemoved == true) && sooners().length > 5) {
      // Reset
      setSooners([{title: title, desc: desc, id: 0, isRemoved: false}])
      return
    }

    setSooners((prevs) => {
      const lastPrevId = prevs[prevs.length - 1]
      
      const res = [...prevs, {title: title, desc: desc, id: lastPrevId ? lastPrevId.id + 1 : 0, isRemoved: false}]

      return res

    })

  };


  const contextPayload: SoonerContextType = {
    show: show,
  };

  return (
    <SoonerContext.Provider value={contextPayload}>
      {/* <Show when={sooners().length > 0}> */}
        <SonnerGroup soonersData={() => sooners()} soonersSetter={setSooners}  />
      {/* </Show> */}
      {props.children}
    </SoonerContext.Provider>
  );
}
