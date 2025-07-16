/* @refresh reload */

import { createContext, createSignal, ParentProps, useContext } from "solid-js";
import Sooner from "~/components/atoms/sooner";
import { AuthContext } from "./auth-context";

interface SoonerType {
    title: string;
    desc: string;
}

interface SoonerContextType {
    show: (title: string, desc: string) => void;
}


export const SoonerContext = createContext<SoonerContextType>()

export function SoonerProvider(props: ParentProps) {

    const [soonerData, setSoonerData] = createSignal<SoonerType|null>(null)
    const [isClosed, setIsClosed] = createSignal<boolean>(true)

    const show = (title: string, desc: string) => {
        setSoonerData({title, desc})
        setIsClosed(false)

        setTimeout(() => {
            setSoonerData(null)
            setIsClosed(true)
        }, 3000)
    }

    const contextPayload: SoonerContextType = {
        show: show
    }


    return (
        <SoonerContext.Provider value={contextPayload}>
            {!isClosed() && 
            <Sooner title={soonerData()?.title} desc={soonerData()?.desc} isClosed={() => isClosed()} setIsClosed={setIsClosed}/>
        }
            {props.children}
        </SoonerContext.Provider>
    )
}
