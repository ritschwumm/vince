import { useState, useCallback, ReactNode, useEffect } from "react";
import { Text, TextInput, FormControl, TreeView, Box, themeGet, IconButton, Octicon } from "@primer/react";
import { PlusIcon, DatabaseIcon } from "@primer/octicons-react";
import { Dialog, PageHeader } from '@primer/react/drafts'

import styled, { css } from "styled-components"
import { Site } from "../../vince";

export const PaneMenu = styled.div`
  position: relative;
  display: flex;
  height: 3rem;
  padding: 0 1rem;
  align-items: center;
  background: ${themeGet("canvas.default")};
  border-top: 1px solid transparent;
  z-index: 5;
`

const loadingStyles = css`
  display: flex;
  justify-content: center;
`

export const PaneWrapper = styled.div`
  display: flex;
  flex-direction: column;
  flex: 1;
`

const Wrapper = styled(PaneWrapper)`
  overflow-x: auto;
  height: 100%;
`

const Menu = styled(PaneMenu)`
  justify-content: space-between;
`


const Header = styled(Text)`
  display: flex;
  align-items: center;
`

type Props = Readonly<{
  children: ReactNode
  className?: string
}>

export const PaneContent = styled.div<Props>`
    display: flex;
    flex-direction: column;
    flex: 1;
    background: ${themeGet("canvas.default")};
    overflow: auto;
  `

const Content = styled(PaneContent) <{
  _loading: boolean
}>`
    display: block;
    overflow: auto;
    ${({ _loading }) => _loading && loadingStyles};
  `




const domainRe = new RegExp("^(?!-)[A-Za-z0-9-]+([-.]{1}[a-z0-9]+)*.[A-Za-z]{2,6}$")

const Sites = () => {
  const [isOpen, setIsOpen] = useState(false);
  const openDialog = useCallback(() => setIsOpen(true), [setIsOpen])
  const closeDialog = useCallback(() => setIsOpen(false), [setIsOpen])
  const [domain, setDomain] = useState<string>("")

  const submitNewSite = useCallback(() => {
    console.log(domain)
    setIsOpen(false)
  }, [domain])
  const [validDomain, setValidDomain] = useState(true)

  useEffect(() => {
    if (domain != "") {
      if (domainRe.test(domain)) {
        setValidDomain(true)
      } else {
        setValidDomain(false)
      }
    }
  }, [domain])

  return (
    <Wrapper>
      <Box paddingX={2}>
        <PageHeader>
          <PageHeader.TitleArea>
            <PageHeader.LeadingAction>
              <DatabaseIcon />
            </PageHeader.LeadingAction>
            <PageHeader.Title>Sites</PageHeader.Title>
            <PageHeader.Actions>
              <IconButton variant="primary" aria-label="New Site" icon={PlusIcon} onClick={openDialog} />
              {isOpen && (
                <Dialog
                  title="Create New Site"
                  footerButtons={
                    [{
                      content: 'Create', onClick: submitNewSite,
                    }]
                  }
                  onClose={closeDialog}
                >
                  <Box>
                    <FormControl>
                      <FormControl.Label>Domain</FormControl.Label>
                      <TextInput
                        monospace
                        block
                        placeholder="vinceanalytics.github.io"
                        onChange={(e) => setDomain(e.currentTarget.value)}
                      />
                      {!validDomain &&
                        <FormControl.Validation id="new-site" variant="error">
                          Domain must be the
                        </FormControl.Validation>}
                    </FormControl>
                  </Box>
                </Dialog>
              )}
            </PageHeader.Actions>
          </PageHeader.TitleArea>
        </PageHeader>
      </Box>
    </Wrapper >
  )
}

export default Sites;