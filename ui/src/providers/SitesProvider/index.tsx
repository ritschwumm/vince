import { createContext, PropsWithChildren, useState, useContext, useEffect } from "react"
import { Site } from "../../vince"
import { useVince } from "../VinceProvider"
import { useLocalStorage, StoreKey } from "../LocalStorageProvider"



type ContextProps = {
    sites: Site[]
    refresh: () => void
}

const defaultValues = {
    sites: [],
    refresh: () => undefined,
}

const SitesContext = createContext<ContextProps>(defaultValues)


export const SitesProvider = ({ children }: PropsWithChildren<{}>) => {
    const [sites, setSites] = useState<Site[]>([])
    const { updateSettings } = useLocalStorage()
    const { vince } = useVince()
    const refresh = () => {
        vince?.listSites({}).then((result) => {
            setSites(result.response.list)
        })
            .catch((e) => {
                console.log(e)
            })
    }
    useEffect(() => {
        if (vince) {
            refresh()
        }
    }, [vince])

    return (
        <SitesContext.Provider
            value={{
                sites,
                refresh,
            }}
        >
            {children}
        </SitesContext.Provider>
    )
}

export const useSites = () => {
    return useContext(SitesContext)
}