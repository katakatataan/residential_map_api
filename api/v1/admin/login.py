from flask_restplus import Namespace, Resource, fields

api = Namespace('login', description='Users login')

class Login(Resource):
    @api.route('/login')
    @api.doc('users_list')
    def get(self):
        '''List all users'''
        return USERS 

USERS = [
    {'id': 'felix', 'name': 'Felix'},
]

class Cat(Resource):
    @api.doc('get_cat')
    @api.marshal_with(cat)
    def get(self, id):
        '''Fetch a cat given its identifier'''
        for cat in CATS:
            if cat['id'] == id:
                return cat
        api.abort(404)
