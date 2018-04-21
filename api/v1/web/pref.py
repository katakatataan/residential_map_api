from flask_restplus import Namespace, Resource, fields

class Prefectures(Resource):
    @api.doc('prefecture list')
    @api.route('/')
    def get(self):
        '''List all cats'''
        return CATS

    @api.route('/<:id>')
    def show(self):
        return CATS
