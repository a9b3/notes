#React
##Components

	var Foo = React.createClass({
		// initializes initial state
		getInitialState() {
			return {
				age: 0
			}
		},
		
		componentDidMount() {
			let fooDiv = React.findDOMNode(this.refs.foo);
		},
		
		// renders dom
		render() {
			var age = this.state.age;
			return (
				<div ref="foo">
					{age}
				</div>
			)
		}
	});
	
defining functions and objects in the argument object, will make it available through 'this', as it would be defined as a property of the object.

Target dom elements by giving them the ref tag. Then you can find them by calling React.findDOMNode(this.refs.foo)


#Meteor
##Reactive Data
	
	// in command line
	> meteor add react
	
	// in react component
	React.createClass({
		// declare ReactMeteorData as a mixin
		mixins: [ReactMeteorData],
		
		// implement the needed function
		// return a reactive data source
		getMeteorData() {
			return {
				user: Meteor.user()
			};
		}
	});
	
After implementing the necessary function 'getMeteorData', the returned object will be available through this.data
Meteor.subscribe can be used inside the function.
<br>
<br>*Only access meteor reactive data source inside getMeteorData*

##Methods
	
	// in command line
	> meteor remove insecure
	
	// in both client and server
	Meteor.methods({
		foo(args) {
			if (!Meteor.userId()) {
				throw new Meteor.Error('not-authorized');
			}
			// else allow
		}
	});
	
	// call secured method now
	Meteor.call('foo', args);
	
Methods provide a way to add security checks for database operations.

	// in command line
	> meteor remove autopublish
	
	// in client
	Meteor.subscribe('tasks');
	
	// in server
	Meteor.publish('tasks', () => {
		return Tasks.find();
	});
	
After subscribe, you can access things in the Tasks collection as usual. If the client is not subscribed the 'tasks' it cannot access Tasks. The publish function can contain some user authentication, or task authorization logic.