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
        var basePath = './src/'
        // check exist src folder
        // if (this.fs.exists('./src')) {
        //     basePath = './src/'
        // }

        var entityName = this.answers.entityName;
        var entityUpCase = entityName.charAt(0).toUpperCase() + entityName.slice(1); // upper case first letter
        var entityLowerCase = entityName.charAt(0).toLowerCase() + entityName.slice(1); // lower case first letter

        var params = {
            entityUpCase: entityUpCase,
            entityLowerCase: entityLowerCase
        }

        this.fs.copyTpl(
            this.templatePath('entity/entity_template.go'),
            this.destinationPath(basePath + 'entity/entity_' + entityLowerCase + '.go'), params
        );

        this.fs.copyTpl(
            this.templatePath('infrastructure/repository/repository_entity.go'),
            this.destinationPath(basePath + 'infrastructure/repository/repository_' + entityLowerCase + '.go'), params
        );

        this.fs.copyTpl(
            this.templatePath('usecase/entity/usecase_entity_interface.go'),
            this.destinationPath(basePath + 'usecase/' + entityLowerCase + '/usecase_' + entityLowerCase + '_interface.go'), params
        );

        this.fs.copyTpl(
            this.templatePath('usecase/entity/usecase_entity_service.go'),
            this.destinationPath(basePath + 'usecase/' + entityLowerCase + '/usecase_' + entityLowerCase + '_service.go'), params
        );

        this.fs.copyTpl(
            this.templatePath('api/handlers/handlers_entity.go'),
            this.destinationPath(basePath + 'api/handlers/handlers_' + entityLowerCase + '.go'), params
        );
    }
}