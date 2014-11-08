/** @jsx React.DOM */

var ApplicationSidebar = React.createClass({
  render: function() {
    return (
      <div>
        <Grid>
          <Row>
            <Col xs={12}>
              <div className='sidebar-header'>PAGES</div>
              <div className='sidebar-nav-container'>
                <SidebarNav style={{marginBottom: 0}}>
                  <SidebarNavItem glyph='icon-fontello-gauge' name='Blank' href='/' />
                  <SidebarNavItem glyph='icon-feather-mail' name={<span>Menu <BLabel className='bg-darkgreen45 fg-white'>3</BLabel></span>}>
                    <SidebarNav>
                      <SidebarNavItem glyph='icon-feather-inbox' name='Inbox' href='#' />
                      <SidebarNavItem glyph='icon-outlined-mail-open' name='Mail' href='#' />
                      <SidebarNavItem glyph='icon-dripicons-message' name='Compose' href='#' />
                    </SidebarNav>
                  </SidebarNavItem>
                </SidebarNav>
              </div>
            </Col>
          </Row>
        </Grid>
      </div>
    );
  }
});

var DummySidebar = React.createClass({
  render: function() {
    return (
      <Grid>
        <Row>
          <Col xs={12}>
            <div className='sidebar-header'>DUMMY SIDEBAR</div>
            <LoremIpsum query='1p' />
          </Col>
        </Row>
      </Grid>
    );
  }
});

var SidebarSection = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <div id='sidebar'>
        <div id='avatar'>
          <Grid>
            <Row className='fg-white'>
              <Col xs={4} collapseRight>
                <img src='/imgs/avatars/avatar0.png' width='40' height='40' />
              </Col>
              <Col xs={8} collapseLeft id='avatar-col'>
                <div style={{top: 23, fontSize: 16, lineHeight: 1, position: 'relative'}}>Anna Sanchez</div>
                <div>
                  <Progress id='demo-progress' value={30} min={0} max={100} color='#ffffff'/>
                  <Link href='#'><Icon id='demo-icon' bundle='fontello' glyph='lock-5' /></Link>
                </div>
              </Col>
            </Row>
          </Grid>
        </div>
        <SidebarControls>
          <SidebarControlBtn bundle='fontello' glyph='docs' key={0} />
          <SidebarControlBtn bundle='fontello' glyph='chat-1' key={1} />
          <SidebarControlBtn bundle='fontello' glyph='chart-pie-2' key={2} />
          <SidebarControlBtn bundle='fontello' glyph='th-list-2' key={3} />
          <SidebarControlBtn bundle='fontello' glyph='bell-5' key={4} />
        </SidebarControls>
        <div id='sidebar-container'>
          <Sidebar key={0} active>
            <ApplicationSidebar />
          </Sidebar>
          <Sidebar key={1}>
            <DummySidebar />
          </Sidebar>
          <Sidebar key={2}>
            <DummySidebar />
          </Sidebar>
          <Sidebar key={3}>
            <DummySidebar />
          </Sidebar>
          <Sidebar key={4}>
            <DummySidebar />
          </Sidebar>
        </div>
      </div>
    );
  }
});

module.exports = SidebarSection;
