/** @jsx React.DOM */

var classSet = React.addons.classSet;

var openState = (!Modernizr.touch) ? (localStorage.getItem('sidebar-open-state') === 'true' ? true : false) : false;
var SidebarMixin = {
  getInitialState: function() {
    return {
      open: openState
    };
  },
  sidebarStateChangeCallback: function(open) {
    if(this.state.open === open) return;
    if(open !== undefined)
      openState = open;
    else
      openState = !this.state.open;
    this.setState({
      open: openState // toggle sidebar
    });
    localStorage.setItem('sidebar-open-state', openState);
  },
  componentWillMount: function() {
    ReactBootstrap.Dispatcher.on('sidebar', this.sidebarStateChangeCallback);
  },
  componentWillUnmount: function() {
    ReactBootstrap.Dispatcher.off('sidebar', this.sidebarStateChangeCallback);
  }
};

var SidebarControlBtn = React.createClass({
  propTypes: {
    key: React.PropTypes.number.isRequired
  },
  getInitialState: function() {
    return {
      active: this.props.active || false
    };
  },
  handleClick: function(e) {
    e.preventDefault();
    e.stopPropagation();
    ReactBootstrap.Dispatcher.emit('sidebar:controlbtn', this.props);
    ReactBootstrap.Dispatcher.emit('sidebar:keychange', this.props.key);
  },
  handleState: function(props) {
    if(props.hasOwnProperty('key')) {
      if(props.key === this.props.key) {
        this.setState({active: true});
      } else {
        this.setState({active: false}); 
      }
    }
  },
  componentDidMount: function() {
    ReactBootstrap.Dispatcher.on('sidebar:controlbtn', this.handleState);
    var scrollToTop = function() {
      if($(window).scrollTop() === 0) return;
      setTimeout(function() {
        $('html, body, #body').scrollTop(0);
        $(window).scrollTop(0);
        scrollToTop();
      }, 15);
    };

    scrollToTop();
  },
  componentWillUnmount: function() {
    ReactBootstrap.Dispatcher.off('sidebar:controlbtn', this.handleState);
  },
  render: function() {
    var classes = classSet({
      'sidebar-control-btn': true,
      'active': this.state.active
    });
    return this.transferPropsTo(
      <li className={classes.trim()} onClick={this.handleClick} tabIndex='-1'>
        <a href='#' tabIndex='-1'>
          <ReactBootstrap.Icon bundle={this.props.bundle} glyph={this.props.glyph}/>
        </a>
      </li>
    );
  }
});

var SidebarControls = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <div className='sidebar-controls-container' dir='ltr'>
        <div className='sidebar-controls' tabIndex='-1'>
          {this.props.children}
        </div>
      </div>
    );
  }
});

var Sidebar = React.createClass({
  propTypes: {
    active: React.PropTypes.bool,
    key: React.PropTypes.number.isRequired
  },
  timer: null,
  getInitialState: function() {
    return {
      left: ((this.props.key * 100) + '%'),
      visibility: (this.props.key === 0) ? 'visible' : 'hidden'
    };
  },
  handleKeyChange: function(key) {
    var newLeft = ((this.props.key*100) - (key*100))+'%';
    var styles = {left: newLeft};
    styles.visibility = 'visible';
    this.setState(styles, function() {
      if(newLeft !== '0%') {
        setTimeout(function() {
          this.setState({visibility: 'hidden'});
        }.bind(this), 300);
      } else if(newLeft === '0%') {
        var height = $(this.refs.innersidebar.getDOMNode()).height();
        if($('html').hasClass('static')) {
          $('#body').css('min-height', height+400);
        } else {
          $('#body').css('min-height', '');
        }
      }
    }.bind(this));
  },
  updateScrollbar: function() {
    if(!Modernizr.touch) {
      $(this.refs.sidebar.getDOMNode()).perfectScrollbar('update');
    }
  },
  repositionScrollbar: function(child_node, top, height) {
    clearTimeout(this.timer);
    var node = $(this.refs.sidebar.getDOMNode());
    var scrollTo = top - node.offset().top + node.scrollTop();
    if(node.find(child_node).length) {
      if(scrollTo > ($(window).height() / 2)) {
        this.timer = setTimeout(function() {
          node.scrollTop(scrollTo - ($(window).height() / 2) + 100);
        }, 15);
      }
    }
    if(!Modernizr.touch)
      this.updateScrollbar()
  },
  destroyScrollbar: function() {
    $(this.refs.sidebar.getDOMNode()).perfectScrollbar('destroy');
  },
  initializeScrollbar: function() {
    $(this.refs.sidebar.getDOMNode()).perfectScrollbar({
      suppressScrollX: true
    });
  },
  componentWillMount: function() {
    ReactBootstrap.Dispatcher.on('sidebar:reinitialize', this.initializeScrollbar);
    ReactBootstrap.Dispatcher.on('sidebar:destroy', this.destroyScrollbar);
    ReactBootstrap.Dispatcher.on('sidebar:update', this.updateScrollbar);
    ReactBootstrap.Dispatcher.on('sidebar:reposition', this.repositionScrollbar);
    ReactBootstrap.Dispatcher.on('sidebar:keychange', this.handleKeyChange);
  },
  componentWillUnmount: function() {
    ReactBootstrap.Dispatcher.off('sidebar:reinitialize', this.initializeScrollbar);
    ReactBootstrap.Dispatcher.off('sidebar:destroy', this.destroyScrollbar);
    ReactBootstrap.Dispatcher.off('sidebar:update', this.updateScrollbar);
    ReactBootstrap.Dispatcher.off('sidebar:reposition', this.repositionScrollbar);
    ReactBootstrap.Dispatcher.off('sidebar:keychange', this.handleKeyChange);
  },
  componentDidMount: function() {
    if(!Modernizr.touch) {
      this.initializeScrollbar();
    }

    if(this.props.active) {
      setTimeout(function() {
        ReactBootstrap.Dispatcher.emit('sidebar:controlbtn', {key: this.props.key});
        ReactBootstrap.Dispatcher.emit('sidebar:keychange', this.props.key);
      }.bind(this), 15);
    }
  },
  render: function() {
    return this.transferPropsTo(
      <div ref='sidebar' className='sidebar' style={{
        left: this.state.left,
        visibility: this.state.visibility,
        'transition': 'all 0.3s ease',
        '-o-transition': 'all 0.3s ease',
        '-ms-transition': 'all 0.3s ease',
        '-moz-transition': 'all 0.3s ease',
        '-webkit-transition': 'all 0.3s ease'
      }}>
        <div ref='innersidebar'>{this.props.children}</div>
      </div>
    );
  }
});

var SidebarNavItem = React.createClass({
  mixins: [RRouter.RoutingContextMixin],
  propTypes: {
    open: React.PropTypes.bool,
    active: React.PropTypes.bool,
    href: React.PropTypes.string,
    autoHeight: React.PropTypes.bool
  },
  getInitialState: function() {
    return {
      open: this.props.open || false,
      active: this.props.active || false,
      toggleOpen: this.props.open || false,
      dir: $('html').attr('dir') === 'ltr' ? 'left' : 'right',
      opposite: $('html').attr('dir') === 'ltr' ? false : true
    };
  },
  emitOpen: function(open) {
    var node = this.refs.node.getDOMNode();
    if(open)
      ReactBootstrap.Dispatcher.emit('sidebar:openstate', node, open);
    else
      ReactBootstrap.Dispatcher.emit('sidebar:openstate', node);
  },
  handleClick: function(e) {
    if(!this.props.href)
      e.preventDefault();
      e.stopPropagation();
    this.emitOpen();
  },
  collapse: function(node, cb) {
    this.setState({
      toggleOpen: false
    }, function() {
      var height = $(node).height();
      $(node).css('height', height).animate({
        height: '45px'
      }, 125, 'easeInOutSine', function() {
        $(node).css('height', '');
        this.setState({
          open: false,
          toggleOpen: false
        }, function() {
          if(cb) cb();
        });
      }.bind(this));
    }.bind(this));
  },
  expand: function(node, cb) {
    this.setState({
      toggleOpen: true
    }, function() {
      var height = $(node).addClass('open').height();
      $(node).removeClass('open');
      $(node).css('height', '45px').animate({
        height: height
      }, 125, 'easeInOutSine', function() {
        $(node).css('height', '');
        this.setState({
          open: true
        }, function() {
          if(cb) cb();
        });
      }.bind(this));      
    }.bind(this));
  },
  handleOpenState: function(child_node, open) {
    if(this.props.children) {
      var node = this.refs.node.getDOMNode();
      if(open) {
        if($(node).find(child_node).length) {
          this.setState({
            open: true,
            toggleOpen: true
          });
        }
        return;
      }
      if($(node).is(child_node)) {
        if(this.state.open) {
          this.collapse(node, function() {
            ReactBootstrap.Dispatcher.emit('sidebar:update');
          });
        } else {
          this.expand(node, function() {
            ReactBootstrap.Dispatcher.emit('sidebar:update');
          });
        }
        return;
      }
      if(!($(node).find(child_node).length)) {
        if(this.state.open)
          this.collapse(node);
      }
    }
  },
  handleLayoutDirChange: function(dir) {
    this.setState({
      dir: dir === 'ltr' ? 'left' : 'right',
      opposite: dir === 'ltr' ? false : true
    });
  },
  componentWillMount: function() {
    ReactBootstrap.Dispatcher.on('layout:dir', this.handleLayoutDirChange);
    ReactBootstrap.Dispatcher.on('sidebar:openstate', this.handleOpenState);
  },
  componentWillUnmount: function() {
    ReactBootstrap.Dispatcher.off('layout:dir', this.handleLayoutDirChange);
    ReactBootstrap.Dispatcher.off('sidebar:openstate', this.handleOpenState);
  },
  componentDidMount: function() {
    var active = (this.getRouting().getPath() === this.props.href) || this.props.active;

    if(active) {
      this.setState({active: true});
      this.emitOpen(true);
      var node = this.refs.node.getDOMNode();
      var height = $(node).height();
      var top = $(node).offset().top;
      ReactBootstrap.Dispatcher.emit('sidebar:reposition', node, top, height);
    }
  },
  render: function() {
    var classes = classSet({
      'open': this.state.open,
      'active': this.state.active
    });
    var toggleClasses = classSet({
      'toggle-button': true,
      'open': this.state.toggleOpen,
      'opposite': this.state.opposite
    });
    var icon=null, toggleButton = null;
    if(this.props.children) {
      toggleButton = <ReactBootstrap.Icon className={toggleClasses.trim()} bundle='fontello' glyph={this.state.dir+'-open-3'} />;
    }
    if(this.props.glyph || this.props.bundle) {
      icon = <ReactBootstrap.Icon bundle={this.props.bundle} glyph={this.props.glyph} />;
    }
    var style = {height: this.props.autoHeight ? 'auto' : ''};
    var componentClass = React.DOM.a;
    if(this.props.href) componentClass = RRouter.Link;
    return this.transferPropsTo(
      <li ref='node' className={classes} style={style} name={null} tabIndex='-1'>
        <componentClass href={this.props.href || '#'} onClick={this.handleClick} name={null} tabIndex='-1' style={style}>
          {icon}
          <span className='name'>{this.props.name}</span>
          {toggleButton}
        </componentClass>
        {this.props.children}
      </li>
    );
  }
});

var SidebarNav = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <ul className='sidebar-nav'>
        {this.props.children}
      </ul>
    );
  }
});

var SidebarBtn = React.createClass({
  handleSidebarStateChange: function(props) {
    if(props['data-id'] === 'sidebar-btn') {
      ReactBootstrap.Dispatcher.emit('sidebar');
    }
  },
  render: function() {
    return this.transferPropsTo(
      <ReactBootstrap.NavContent className='pull-left visible-xs-inline-block'>
        <ReactBootstrap.Nav onItemSelect={this.handleSidebarStateChange}>
          <ReactBootstrap.NavItem data-id='sidebar-btn' className='sidebar-btn' href='/'>
            <ReactBootstrap.Icon bundle='fontello' glyph='th-list-5' />
          </ReactBootstrap.NavItem>
        </ReactBootstrap.Nav>
      </ReactBootstrap.NavContent>
    );
  }
});

module.exports = {
  Sidebar: Sidebar,
  SidebarNav: SidebarNav,
  SidebarBtn: SidebarBtn,
  SidebarMixin: SidebarMixin,
  SidebarNavItem: SidebarNavItem,
  SidebarControls: SidebarControls,
  SidebarControlBtn: SidebarControlBtn
};
