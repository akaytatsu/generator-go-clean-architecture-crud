import Generator from "yeoman-generator";

export default class extends Generator {

    initializing() {
        this.props = {};
    }

    async prompting() {
        this.answers = await this.prompt([{
            type: 'input',
            name: 'entityName',
            message: 'name of entity without word "Entity" (i.e.: User):',
        }]);
    }

    writing() {
        var basePath = './app/'

        var entityName = this.answers.entityName;
        var entityUpCase = entityName.charAt(0).toUpperCase() + entityName.slice(1); // upper case first letter
        var entityLowerCase = entityName.charAt(0).toLowerCase() + entityName.slice(1); // lower case first letter

        var params = {
            entityUpCase: entityUpCase,
            entityLowerCase: entityLowerCase
        }

        this.fs.copyTpl(
            this.templatePath('models/entity.tsx'),
            this.destinationPath(basePath + 'models/' + entityLowerCase + '.model.ts'), params
        );

        this.fs.copyTpl(
            this.templatePath('routes/_app.entity/columns.tsx'),
            this.destinationPath(basePath + 'routes/_app.' + entityLowerCase + '/columns.tsx'), params
        );

        this.fs.copyTpl(
            this.templatePath('routes/_app.entity/addOrEdit.tsx'),
            this.destinationPath(basePath + 'routes/_app.' + entityLowerCase + '/addOrEdit.tsx'), params
        );

        this.fs.copyTpl(
            this.templatePath('routes/_app.entity/route.tsx'),
            this.destinationPath(basePath + 'routes/_app.' + entityLowerCase + '/route.tsx'), params
        );

        this.fs.copyTpl(
            this.templatePath('services/entity.service.ts'),
            this.destinationPath(basePath + 'services/' + entityLowerCase + '.service.ts'), params
        );
    }
}